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
// Wrappers without a discriminator are left alone — their variants are
// usually structurally distinct enough for the naive matcher to work.
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

	fixed, err := walkAndFix(*clientDir, discriminators)
	if err != nil {
		log.Fatalf("walk %s: %v", *clientDir, err)
	}
	log.Printf("done — %d file(s) rewritten", fixed)
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

func walkAndFix(dir string, discs map[string]Discriminator) (int, error) {
	count := 0
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		fixed, err := tryFixFile(path, discs)
		if err != nil {
			return err
		}
		if fixed {
			count++
		}
		return nil
	})
	return count, err
}

// variant captures one field of a oneOf wrapper struct: the field name as it
// appears in the Go code, and the pointed-to type name.
type variant struct {
	fieldName string
	typeName  string
}

func tryFixFile(path string, discs map[string]Discriminator) (bool, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return false, fmt.Errorf("%s: parse: %w", path, err)
	}

	typeName, variants, ok := findOneOfWrapper(file)
	if !ok {
		return false, nil
	}
	disc, ok := discs[typeName]
	if !ok {
		return false, nil
	}
	if !hasNaiveUnmarshal(file, typeName) {
		return false, nil
	}

	if err := rewriteFile(fset, file, path, typeName, variants, disc); err != nil {
		return false, fmt.Errorf("%s: rewrite: %w", path, err)
	}
	log.Printf("rewrote UnmarshalJSON for %s in %s", typeName, path)
	return true, nil
}

// findOneOfWrapper scans the file for a top-level struct whose every field is
// a pointer (the openapi-generator convention for oneOf wrappers). Returns
// the struct name and the list of variants. Skips structs with mixed shapes.
func findOneOfWrapper(file *ast.File) (typeName string, variants []variant, ok bool) {
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
			st, ok := ts.Type.(*ast.StructType)
			if !ok || st.Fields == nil {
				continue
			}
			vars, ok := extractPointerVariants(st)
			if !ok {
				continue
			}
			return ts.Name.Name, vars, true
		}
	}
	return "", nil, false
}

func extractPointerVariants(st *ast.StructType) ([]variant, bool) {
	var variants []variant
	for _, f := range st.Fields.List {
		star, ok := f.Type.(*ast.StarExpr)
		if !ok {
			return nil, false
		}
		ident, ok := star.X.(*ast.Ident)
		if !ok {
			return nil, false
		}
		if len(f.Names) != 1 {
			// Embedded field or grouped declaration; skip — not a oneOf wrapper.
			return nil, false
		}
		variants = append(variants, variant{
			fieldName: f.Names[0].Name,
			typeName:  ident.Name,
		})
	}
	if len(variants) < 2 {
		return nil, false
	}
	return variants, true
}

// hasNaiveUnmarshal returns true if the file's UnmarshalJSON contains the
// exact error string the buggy generator template emits. We match on that
// rather than reconstructing the whole pattern — it's the closest thing to a
// signature this template has, and it costs us nothing if we get a false
// negative (we just leave the file alone).
func hasNaiveUnmarshal(file *ast.File, typeName string) bool {
	needle := fmt.Sprintf("data matches more than one schema in oneOf(%s)", typeName)
	for _, decl := range file.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok || fd.Name.Name != "UnmarshalJSON" || fd.Recv == nil {
			continue
		}
		if !receiverIs(fd.Recv, typeName) {
			continue
		}
		var buf bytes.Buffer
		if err := format.Node(&buf, token.NewFileSet(), fd); err == nil {
			if strings.Contains(buf.String(), needle) {
				return true
			}
		}
	}
	return false
}

func rewriteFile(fset *token.FileSet, file *ast.File, path, typeName string, variants []variant, disc Discriminator) error {
	// Locate the old UnmarshalJSON we're replacing so we can splice it out of
	// the source bytes by byte offset. Grafting an AST node parsed under a
	// different FileSet causes go/format to interleave the old body's
	// comments at stale positions; text splice + reparse is cleaner.
	var oldFd *ast.FuncDecl
	for _, d := range file.Decls {
		fd, ok := d.(*ast.FuncDecl)
		if !ok || fd.Name.Name != "UnmarshalJSON" || fd.Recv == nil {
			continue
		}
		if !receiverIs(fd.Recv, typeName) {
			continue
		}
		oldFd = fd
		break
	}
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
	b.WriteString("\t// Postprocessed by scripts/fix-oneof-decoder: dispatch by the OpenAPI\n")
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
