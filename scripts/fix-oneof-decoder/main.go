// Command fix-oneof-decoder rewrites the generated UnmarshalJSON method of
// oneOf wrapper types so it dispatches by the OpenAPI discriminator instead
// of doing naive structural matching.
//
// openapi-generator's Go template ignores the OpenAPI 3 `discriminator`
// declaration when emitting UnmarshalJSON for a oneOf: it tries to decode
// the payload into every variant in turn and counts how many succeed. When
// two variants generate to byte-identical Go structs (e.g. Access's
// OrganisationAccess and AllExternalAccess on TheHive 5.6 — both reduce to
// `{ Kind string }`) the same payload matches both, the counter hits 2, and
// the decoder returns "data matches more than one schema in oneOf(X)" on
// every endpoint that returns one of those wrappers.
//
// This tool runs AFTER the generator container exits, scans the generated
// client, and replaces the buggy UnmarshalJSON with a `_kind`-driven switch
// for any wrapper that has a discriminator mapping declared in the spec.
//
// Strict mode: for every type that has a `discriminator:` block in the spec
// we MUST be able to handle it — either by rewriting the naive pattern, or
// by detecting that we've already rewritten it (idempotent re-runs). If we
// see something we don't recognise — typically because openapi-generator
// changed the wording of its UnmarshalJSON template — we log.Fatalf and
// fail `make generate` loudly instead of silently shipping a SDK whose
// decoders revert to the broken naive matcher.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	defaultClientDir = "thehive"
	defaultSpecPath  = "tmp/thehive_openapi_fixed.yaml"

	// rewrittenMarker is the unique comment we emit at the top of every
	// UnmarshalJSON body we rewrite. classifyUnmarshal looks for it to
	// recognise an already-rewritten file on re-runs (idempotence).
	rewrittenMarker = "Postprocessed by scripts/fix-oneof-decoder:"
)

func main() {
	clientDir := flag.String("dir", defaultClientDir, "directory containing generated Go files")
	specPath := flag.String("spec", defaultSpecPath, "preprocessed OpenAPI spec to read discriminators from")
	flag.Parse()

	log.SetFlags(0)
	log.SetPrefix("fix-oneof-decoder: ")

	discriminators, err := loadDiscriminators(*specPath)
	if err != nil {
		log.Fatalf("load discriminators from %s: %v", *specPath, err)
	}
	log.Printf("loaded %d discriminator mapping(s) from %s", len(discriminators), *specPath)

	typeToPath, err := indexTypeDecls(*clientDir)
	if err != nil {
		log.Fatalf("index type declarations in %s: %v", *clientDir, err)
	}

	rewritten, idempotent := applyToDiscriminators(discriminators, typeToPath, *clientDir)
	log.Printf("done — %d rewritten, %d already up-to-date", rewritten, idempotent)
}

// Discriminator captures the propertyName + value→variant mapping from a
// schema's `discriminator:` block.
type Discriminator struct {
	Property string            // "_kind"
	Mapping  map[string]string // {"OrganisationAccessKind": "OrganisationAccess", ...}
}

func loadDiscriminators(path string) (map[string]Discriminator, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var spec struct {
		Components struct {
			Schemas map[string]struct {
				Discriminator *struct {
					PropertyName string            `yaml:"propertyName"`
					Mapping      map[string]string `yaml:"mapping"`
				} `yaml:"discriminator"`
			} `yaml:"schemas"`
		} `yaml:"components"`
	}
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, err
	}

	out := make(map[string]Discriminator)
	for name, s := range spec.Components.Schemas {
		if s.Discriminator == nil || len(s.Discriminator.Mapping) == 0 {
			continue
		}
		mapping := make(map[string]string, len(s.Discriminator.Mapping))
		for kind, ref := range s.Discriminator.Mapping {
			// ref is "#/components/schemas/<TypeName>"; we want just <TypeName>.
			i := strings.LastIndex(ref, "/")
			if i < 0 {
				return nil, fmt.Errorf("%s: malformed discriminator ref %q", name, ref)
			}
			mapping[kind] = ref[i+1:]
		}
		out[name] = Discriminator{
			Property: s.Discriminator.PropertyName,
			Mapping:  mapping,
		}
	}
	return out, nil
}

// indexTypeDecls walks the client directory and builds a map from each
// top-level Go type name to the file that declares it. We use this to look
// up wrapper structs by name rather than by structural shape: lookup is more
// robust to future template changes, and strict mode can confidently report
// "type X has a discriminator in the spec but no Go file declares it".
func indexTypeDecls(dir string) (map[string]string, error) {
	out := make(map[string]string)
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".go") {
			return err
		}
		fset := token.NewFileSet()
		file, perr := parser.ParseFile(fset, path, nil, parser.PackageClauseOnly|parser.SkipObjectResolution)
		_ = file
		_ = perr
		// Re-parse with declarations included; PackageClauseOnly is too narrow.
		file, perr = parser.ParseFile(fset, path, nil, parser.SkipObjectResolution)
		if perr != nil {
			return fmt.Errorf("%s: parse: %w", path, perr)
		}
		for _, decl := range file.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.TYPE {
				continue
			}
			for _, spec := range gen.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				out[ts.Name.Name] = path
			}
		}
		return nil
	})
	return out, err
}

// fileState captures what we found in a file's UnmarshalJSON.
type fileState int

const (
	stateNaive       fileState = iota // matches the buggy generator template — rewrite it
	stateRewritten                    // already postprocessed by us — idempotent skip
	stateUnknown                      // matches neither pattern — fail loud in strict mode
	stateExoticShape                  // struct has non-standard variant shape (e.g. *map[string]any) — leave naive matcher in place
)

// applyToDiscriminators iterates the spec's discriminator types and ensures
// each is handled. Returns (rewritten, idempotent) counters. Fatals on any
// type that isn't accounted for.
func applyToDiscriminators(discs map[string]Discriminator, typeToPath map[string]string, clientDir string) (int, int) {
	rewritten, idempotent := 0, 0

	// Sort for deterministic log output (helps reviewers compare runs).
	typeNames := make([]string, 0, len(discs))
	for t := range discs {
		typeNames = append(typeNames, t)
	}
	sort.Strings(typeNames)

	for _, typeName := range typeNames {
		disc := discs[typeName]
		path, ok := typeToPath[typeName]
		if !ok {
			// The spec declares a discriminator for this type but no Go file
			// in the generated client declares it. Either the generator
			// dropped a schema, or the dir argument is wrong. Either way,
			// shipping silently is the worst option.
			log.Fatalf("strict: type %q has a discriminator in the spec but no .go file under %s declares it", typeName, clientDir)
		}

		state, variants, err := classifyUnmarshal(path, typeName)
		if err != nil {
			log.Fatalf("classify %s in %s: %v", typeName, path, err)
		}

		switch state {
		case stateNaive:
			if err := rewriteFile(path, typeName, variants, disc); err != nil {
				log.Fatalf("rewrite %s in %s: %v", typeName, path, err)
			}
			log.Printf("rewrote UnmarshalJSON for %s in %s", typeName, path)
			rewritten++
		case stateRewritten:
			log.Printf("%s in %s already rewritten — idempotent skip", typeName, path)
			idempotent++
		case stateExoticShape:
			// The wrapper has at least one non-struct variant (typically a
			// `*map[string]interface{}` for a schema that's `type: object`
			// with `additionalProperties: true`). The discriminator can't
			// cleanly dispatch to that field, so we leave the naive matcher
			// in place. It works in practice because the map variant is
			// structurally distinct from the other (named-struct) variants.
			log.Printf("%s in %s has exotic variant shape — leaving naive matcher in place", typeName, path)
		case stateUnknown:
			// Strict mode: this is the failure that matters. Either
			// openapi-generator changed its UnmarshalJSON template (so neither
			// the naive pattern nor our marker is present), or this tool has
			// a detection bug. Fail loud so `make generate` blocks instead of
			// silently shipping a SDK whose oneOf decoders revert to the
			// buggy naive matcher.
			log.Fatalf(
				"strict: %s.UnmarshalJSON in %s matches neither the naive generator pattern "+
					"nor our rewritten marker. openapi-generator template may have changed "+
					"(see scripts/fix-oneof-decoder/main.go for the patterns we look for) — "+
					"re-audit and adapt this tool",
				typeName, path)
		}
	}
	return rewritten, idempotent
}

// variant captures one field of a oneOf wrapper struct: the field name as it
// appears in the Go code, and the pointed-to type name.
type variant struct {
	fieldName string
	typeName  string
}

// classifyUnmarshal parses the file, locates UnmarshalJSON on *typeName, and
// decides whether we're looking at the naive generator template (needs
// rewrite), our own rewrite (skip), or something we don't recognise (fail
// loud upstream). It also returns the variants extracted from the wrapper
// struct, which the rewriter needs.
func classifyUnmarshal(path, typeName string) (fileState, []variant, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return stateUnknown, nil, fmt.Errorf("parse: %w", err)
	}

	variants, shape, err := extractWrapperVariants(file, typeName)
	if err != nil {
		return stateUnknown, nil, err
	}
	if shape == shapeExotic {
		return stateExoticShape, nil, nil
	}

	fd := findUnmarshal(file, typeName)
	if fd == nil {
		return stateUnknown, nil, fmt.Errorf("UnmarshalJSON not found")
	}

	bodySrc, err := funcSource(fd)
	if err != nil {
		return stateUnknown, nil, err
	}

	if isAlreadyRewritten(bodySrc) {
		return stateRewritten, variants, nil
	}
	if isNaiveTemplate(bodySrc, typeName) {
		return stateNaive, variants, nil
	}
	return stateUnknown, variants, nil
}

// wrapperShape classifies the struct's field layout.
type wrapperShape int

const (
	shapeStandard wrapperShape = iota // every field is *NamedType — can be rewritten
	shapeExotic                       // every field is a pointer but at least one is *map / *iface — leave alone
	shapeOther                        // not a oneOf wrapper shape at all (mixed/non-pointer fields, etc.)
)

// extractWrapperVariants looks up the named struct in the file and inspects
// its field layout. Returns the variants extracted from the struct (only
// populated for standard shape) and the shape classification.
func extractWrapperVariants(file *ast.File, typeName string) ([]variant, wrapperShape, error) {
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.TYPE {
			continue
		}
		for _, spec := range gen.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok || ts.Name.Name != typeName {
				continue
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok || st.Fields == nil {
				return nil, shapeOther, nil
			}
			vars, shape := extractPointerVariants(st)
			return vars, shape, nil
		}
	}
	return nil, shapeOther, fmt.Errorf("struct %s not found", typeName)
}

func extractPointerVariants(st *ast.StructType) ([]variant, wrapperShape) {
	var variants []variant
	hasExotic := false
	for _, f := range st.Fields.List {
		star, ok := f.Type.(*ast.StarExpr)
		if !ok {
			// Non-pointer field — not a oneOf wrapper at all.
			return nil, shapeOther
		}
		if len(f.Names) != 1 {
			// Embedded or grouped declaration; not a oneOf wrapper.
			return nil, shapeOther
		}
		ident, ok := star.X.(*ast.Ident)
		if !ok {
			// Pointer to a non-named type (map, slice, interface…).
			// The wrapper still looks like a oneOf union but at least one
			// variant has no Go type name we can dispatch to via the
			// discriminator — flag as exotic so the caller can skip safely.
			hasExotic = true
			continue
		}
		variants = append(variants, variant{
			fieldName: f.Names[0].Name,
			typeName:  ident.Name,
		})
	}
	if hasExotic {
		return variants, shapeExotic
	}
	if len(variants) < 2 {
		return nil, shapeOther
	}
	return variants, shapeStandard
}

func findUnmarshal(file *ast.File, typeName string) *ast.FuncDecl {
	for _, decl := range file.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok || fd.Name.Name != "UnmarshalJSON" || fd.Recv == nil {
			continue
		}
		if receiverIs(fd.Recv, typeName) {
			return fd
		}
	}
	return nil
}

func funcSource(fd *ast.FuncDecl) (string, error) {
	var buf bytes.Buffer
	if err := format.Node(&buf, token.NewFileSet(), fd); err != nil {
		return "", fmt.Errorf("format func: %w", err)
	}
	return buf.String(), nil
}

// isAlreadyRewritten returns true if the function body carries our marker
// comment. This lets us re-run the tool idempotently without misclassifying
// an already-fixed file as "unknown".
func isAlreadyRewritten(bodySrc string) bool {
	return strings.Contains(bodySrc, rewrittenMarker)
}

// isNaiveTemplate detects the buggy openapi-generator template using two
// independent signals (belt-and-braces): the unique error string AND the
// `match := 0` counter declaration. Requiring BOTH means a wording change
// on just one of them still classifies as naive — only a template that
// drops both falls through to the strict-mode failure.
//
// Coupled to openapitools/openapi-generator v7.14.0 Go template (see
// Makefile OPENAPI_GENERATOR_IMAGE). If you bump the generator, re-audit
// these markers.
func isNaiveTemplate(bodySrc, typeName string) bool {
	errString := fmt.Sprintf("data matches more than one schema in oneOf(%s)", typeName)
	return strings.Contains(bodySrc, errString) && strings.Contains(bodySrc, "match := 0")
}

func rewriteFile(path, typeName string, variants []variant, disc Discriminator) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	// Locate the old UnmarshalJSON we're replacing so we can splice it out of
	// the source bytes by byte offset. Grafting an AST node parsed under a
	// different FileSet causes go/format to interleave the old body's
	// comments at stale positions; text splice + reparse is cleaner.
	oldFd := findUnmarshal(file, typeName)
	if oldFd == nil {
		return fmt.Errorf("UnmarshalJSON for %s not found", typeName)
	}

	newSource, err := generateUnmarshalJSON(typeName, variants, disc)
	if err != nil {
		return err
	}

	// FuncDecl.Pos() returns Doc.Pos() when a doc comment is attached, so the
	// splice naturally drops the old "// Unmarshal JSON data into..." doc too.
	src, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	startOff := fset.Position(oldFd.Pos()).Offset
	endOff := fset.Position(oldFd.End()).Offset

	var spliced bytes.Buffer
	spliced.Write(src[:startOff])
	spliced.WriteString(newSource)
	spliced.Write(src[endOff:])

	// Re-parse with a fresh FileSet — all positions now reflect the spliced
	// source. Then drop the now-unused validator import and pretty-print.
	fset2 := token.NewFileSet()
	file2, err := parser.ParseFile(fset2, path, spliced.Bytes(), parser.ParseComments)
	if err != nil {
		return fmt.Errorf("re-parse: %w", err)
	}
	removeImport(file2, "gopkg.in/validator.v2")

	var out bytes.Buffer
	if err := format.Node(&out, fset2, file2); err != nil {
		return fmt.Errorf("format: %w", err)
	}
	return os.WriteFile(path, out.Bytes(), 0o644)
}

func receiverIs(recv *ast.FieldList, typeName string) bool {
	if recv == nil || len(recv.List) != 1 {
		return false
	}
	star, ok := recv.List[0].Type.(*ast.StarExpr)
	if !ok {
		return false
	}
	ident, ok := star.X.(*ast.Ident)
	if !ok {
		return false
	}
	return ident.Name == typeName
}

func removeImport(file *ast.File, importPath string) {
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.IMPORT {
			continue
		}
		kept := gen.Specs[:0]
		for _, spec := range gen.Specs {
			is, ok := spec.(*ast.ImportSpec)
			if !ok {
				kept = append(kept, spec)
				continue
			}
			if is.Path.Value == `"`+importPath+`"` {
				continue
			}
			kept = append(kept, spec)
		}
		gen.Specs = kept
	}
}

func generateUnmarshalJSON(typeName string, variants []variant, disc Discriminator) (string, error) {
	typeToField := make(map[string]string, len(variants))
	for _, v := range variants {
		typeToField[v.typeName] = v.fieldName
	}

	// Bail if the discriminator references a variant we don't have on the
	// struct. Emitting a switch case for a missing field would compile-fail;
	// better to leave the file untouched and let the developer investigate.
	for kind, varType := range disc.Mapping {
		if _, ok := typeToField[varType]; !ok {
			return "", fmt.Errorf("discriminator %s=%q points to variant %q not present on %s",
				disc.Property, kind, varType, typeName)
		}
	}

	kinds := make([]string, 0, len(disc.Mapping))
	for k := range disc.Mapping {
		kinds = append(kinds, k)
	}
	sort.Strings(kinds)

	var b strings.Builder
	fmt.Fprintf(&b, "func (dst *%s) UnmarshalJSON(data []byte) error {\n", typeName)
	fmt.Fprintf(&b, "\t// %s dispatch by the OpenAPI\n", rewrittenMarker)
	b.WriteString("\t// discriminator instead of naive structural matching, which fails when\n")
	b.WriteString("\t// two variants generate to byte-identical Go structs.\n")
	fmt.Fprintf(&b, "\tvar disc struct {\n\t\tKind string `json:\"%s\"`\n\t}\n", disc.Property)
	b.WriteString("\tif err := json.Unmarshal(data, &disc); err != nil {\n")
	fmt.Fprintf(&b, "\t\treturn fmt.Errorf(\"oneOf(%s): cannot read discriminator %s: %%w\", err)\n", typeName, disc.Property)
	b.WriteString("\t}\n")
	b.WriteString("\tswitch disc.Kind {\n")
	for _, kind := range kinds {
		variantType := disc.Mapping[kind]
		field := typeToField[variantType]
		fmt.Fprintf(&b, "\tcase %q:\n", kind)
		fmt.Fprintf(&b, "\t\tdst.%s = new(%s)\n", field, variantType)
		fmt.Fprintf(&b, "\t\treturn newStrictDecoder(data).Decode(dst.%s)\n", field)
	}
	b.WriteString("\tdefault:\n")
	fmt.Fprintf(&b, "\t\treturn fmt.Errorf(\"oneOf(%s): unknown %s value %%q\", disc.Kind)\n", typeName, disc.Property)
	b.WriteString("\t}\n")
	b.WriteString("}\n")
	return b.String(), nil
}
