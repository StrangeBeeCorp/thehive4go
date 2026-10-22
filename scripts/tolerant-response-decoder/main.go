// Command tolerant-response-decoder rewrites the generated UnmarshalJSON of
// every model TheHive sends back, so that a newer TheHive does not break the
// client.
//
// openapi-generator's Go template decodes a model that has required
// properties with json.Decoder.DisallowUnknownFields. TheHive adds fields to
// its responses in minor versions (5.8 added `extraData` to actions,
// `timeToHandle` to tasks; 6.0 adds `requiresImpact` to case statuses), so a
// client generated from the 5.6 spec rejected every 5.8 responder action
// although TheHive had created it. A client must ignore what it does not know
// in a response: this tool removes DisallowUnknownFields from every model
// reachable from a response in the spec. Top-level request models (Input*)
// stay strict, since a consumer may decode a caller's input into them; a type
// shared between a request and a response (Access, in InputManageCaseAccess,
// for one) becomes lenient. That is the accepted trade-off.
//
// It also rewrites the naive oneOf decoder of wrappers whose variants are all
// arrays of primitives (OutputCustomFieldOptions: []float64 | []int32 |
// []string). There, the template counts how many variants decode and fails
// when more than one does, which `[]` (every variant) and `[1, 2]` (float and
// int) always do. The rewrite takes the first variant that decodes, trying the
// most specific type first (integers, floats, booleans, strings): `[1, 2]` is
// an []int32, `[1.5]` an []float64 and `[]` the first of the order.
//
// The variants of the other naive oneOf wrappers keep DisallowUnknownFields:
// that strictness is what lets the naive matcher tell them apart. Only one
// kind of wrapper is spared: one that no request refers to, whose variants are
// all model types that check their own required properties and that a
// first-match decoder can tell apart, the variant with the most required
// properties first (OutputLicenseCurrent). Its decoder is rewritten and its
// variants become lenient. Filter, which requests build, and
// OutputInvokeFunction, whose map variant matches anything, keep theirs.
//
// The same goes for enums. The generated UnmarshalJSON of an enum accepts only
// the values of the spec the client was generated from and fails the whole
// answer on any other, so TheHive adding a value to an enum in a minor version
// broke it (the 6.0 permission names, for one). The enums reachable from a
// response keep the value they do not list. An enum that only requests refer to
// stays strict, like the request models do; one shared with a request becomes
// lenient, like a shared model does.
//
// Strict mode: the tool fails when a response model has an UnmarshalJSON it
// does not recognise, when a schema has no Go type, when a model a
// discriminated oneOf dispatches to has no UnmarshalJSON of its own (the
// dispatch would decode it with DisallowUnknownFields again), when an enum
// decoder has a shape it does not recognise, or when, after it ran, a response
// model still disallows unknown fields without being on the list of naive oneOf
// variants or a response enum still rejects a value it does not list.
// Idempotent: every rewrite leaves a marker comment.
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
	"reflect"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	defaultClientDir = "thehive"
	defaultSpecPath  = "thehive/api/openapi.yaml"

	schemaRefPrefix = "#/components/schemas/"

	// lenientMarker replaces the DisallowUnknownFields call it removes.
	lenientMarker = "Postprocessed by scripts/tolerant-response-decoder: unknown fields are ignored"
	// firstMatchMarker opens the oneOf decoders this tool rewrites.
	firstMatchMarker = "Postprocessed by scripts/tolerant-response-decoder: first match wins"

	// enumMarker replaces the validation of a response enum.
	enumMarker = "Postprocessed by scripts/tolerant-response-decoder: unknown enum values are kept"

	// dispatchMarker opens the discriminated oneOf decoders of fix-oneof-decoder.
	dispatchMarker = "Postprocessed by scripts/fix-oneof-decoder:"

	disallowCall = "decoder.DisallowUnknownFields()"
)

func main() {
	clientDir := flag.String("dir", defaultClientDir, "directory containing generated Go files")
	specPath := flag.String("spec", defaultSpecPath, "OpenAPI spec the client was generated from")
	flag.Parse()

	log.SetFlags(0)
	log.SetPrefix("tolerant-response-decoder: ")

	spec, err := loadSpec(*specPath)
	if err != nil {
		log.Fatalf("load spec %s: %v", *specPath, err)
	}

	types, err := indexTypes(*clientDir)
	if err != nil {
		log.Fatalf("index types in %s: %v", *clientDir, err)
	}

	reachable := spec.responseSchemas()
	log.Printf("%d schema(s) reachable from a response", len(reachable))

	responseTypes := make([]string, 0, len(reachable))
	for _, schema := range reachable {
		goType, ok := types.lookup(schema)
		if !ok {
			if spec.isObject(schema) {
				log.Fatalf("strict: schema %q is an object reachable from a response, but no Go type under %s declares it", schema, *clientDir)
			}
			continue
		}
		responseTypes = append(responseTypes, goType)
	}
	sort.Strings(responseTypes)

	requestTypes := map[string]bool{}
	for _, schema := range spec.requestSchemas() {
		if goType, ok := types.lookup(schema); ok {
			requestTypes[goType] = true
		}
	}

	strictVariants := map[string]string{}
	firstMatch, spared := 0, 0
	for _, goType := range responseTypes {
		w, err := inspectWrapper(types.files[goType], goType)
		if err != nil {
			log.Fatalf("inspect %s: %v", goType, err)
		}
		if w == nil {
			continue
		}
		switch {
		case w.primitiveArrays:
			changed, err := rewriteFirstMatch(types.files[goType], goType, bySpecificity(w.fields), primitiveReason)
			if err != nil {
				log.Fatalf("rewrite %s: %v", goType, err)
			}
			if changed {
				log.Printf("%s: rewrote the naive oneOf decoder to first match wins, most specific type first", goType)
			} else {
				log.Printf("%s: first-match decoder already in place — idempotent skip", goType)
			}
			firstMatch++
		case len(w.namedVariants) == 0:
			// Nothing to keep strict: its variants are slices or maps.
		case requestTypes[goType]:
			log.Printf("%s: variants kept strict, a request refers to it", goType)
			for _, v := range w.namedVariants {
				strictVariants[v] = goType
			}
		default:
			ordered, why := orderVariants(w, types.files)
			if why != "" {
				if w.rewritten {
					log.Fatalf("strict: %s has a first-match decoder, but its variants cannot be ordered any more: %s", goType, why)
				}
				log.Printf("%s: variants kept strict, %s", goType, why)
				for _, v := range w.namedVariants {
					strictVariants[v] = goType
				}
				continue
			}
			changed, err := rewriteFirstMatch(types.files[goType], goType, ordered, namedReason)
			if err != nil {
				log.Fatalf("rewrite %s: %v", goType, err)
			}
			if changed {
				log.Printf("%s: variants made lenient, decoder rewritten to first match wins, most required properties first", goType)
			} else {
				log.Printf("%s: variants lenient, first-match decoder already in place — idempotent skip", goType)
			}
			firstMatch++
			spared += len(w.namedVariants)
		}
	}

	lenient, already := 0, 0
	for _, goType := range responseTypes {
		if wrapper, kept := strictVariants[goType]; kept {
			log.Printf("%s: kept strict, a variant of the naive oneOf %s", goType, wrapper)
			continue
		}
		changed, err := dropDisallow(types.files[goType], goType)
		if err != nil {
			log.Fatalf("%s: %v", goType, err)
		}
		switch changed {
		case resultRewritten:
			lenient++
		case resultAlready:
			already++
		}
	}

	enumsLenient, enumsAlready, enumsShared := 0, 0, 0
	for _, goType := range responseTypes {
		changed, err := relaxEnum(types.files[goType], goType)
		if err != nil {
			log.Fatalf("%s: %v", goType, err)
		}
		if changed == resultNone {
			continue
		}
		if requestTypes[goType] {
			enumsShared++
			log.Printf("%s: enum also used by a request, lenient too", goType)
		}
		switch changed {
		case resultRewritten:
			enumsLenient++
		case resultAlready:
			enumsAlready++
		}
	}

	if err := verify(responseTypes, strictVariants, types.files); err != nil {
		log.Fatalf("strict: %v", err)
	}
	log.Printf("done — %d decoder(s) made lenient, %d already lenient, %d oneOf decoder(s) first-match, %d naive oneOf variant(s) kept strict, %d naive oneOf variant(s) lenient, %d enum(s) made lenient, %d already lenient (%d shared with a request)",
		lenient, already, firstMatch, len(strictVariants), spared, enumsLenient, enumsAlready, enumsShared)
}

// spec is the part of the OpenAPI document this tool reads.
type spec struct {
	Paths      map[string]map[string]yaml.Node `yaml:"paths"`
	Components struct {
		Schemas map[string]yaml.Node `yaml:"schemas"`
	} `yaml:"components"`
}

func loadSpec(path string) (*spec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var s spec
	if err := yaml.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	if len(s.Paths) == 0 || len(s.Components.Schemas) == 0 {
		return nil, fmt.Errorf("no paths or no component schemas: is %s an OpenAPI document?", path)
	}
	return &s, nil
}

// responseSchemas lists, sorted, every component schema a response body
// refers to, directly or through other schemas.
func (s *spec) responseSchemas() []string {
	return s.schemasUnder("responses")
}

// requestSchemas lists, sorted, every component schema a request refers to
// (its body or its parameters), directly or through other schemas.
func (s *spec) requestSchemas() []string {
	return s.schemasUnder("requestBody", "parameters")
}

// schemasUnder lists, sorted, every component schema that the given parts of
// the operations refer to, directly or through other schemas. A path may also
// declare `parameters` for all its operations.
func (s *spec) schemasUnder(parts ...string) []string {
	seen := map[string]bool{}
	var queue []string
	add := func(node *yaml.Node) {
		for _, name := range refs(node) {
			if !seen[name] {
				seen[name] = true
				queue = append(queue, name)
			}
		}
	}

	for _, ops := range s.Paths {
		for method, op := range ops {
			if method == "parameters" {
				for _, part := range parts {
					if part == "parameters" {
						add(&op)
					}
				}
				continue
			}
			if op.Kind != yaml.MappingNode {
				continue
			}
			for _, part := range parts {
				if node := child(&op, part); node != nil {
					add(node)
				}
			}
		}
	}

	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]
		if node, ok := s.Components.Schemas[name]; ok {
			add(&node)
		}
	}

	out := make([]string, 0, len(seen))
	for name := range seen {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}

// isObject reports whether the schema would generate a Go type: an object with
// properties, or a composition of schemas. A pure map (`additionalProperties`
// alone, like Map_Boolean) generates none: the models that use it declare a
// map[string]bool inline.
func (s *spec) isObject(name string) bool {
	node, ok := s.Components.Schemas[name]
	if !ok {
		return false
	}
	for _, key := range []string{"properties", "oneOf", "allOf", "anyOf"} {
		if child(&node, key) != nil {
			return true
		}
	}
	return false
}

func child(node *yaml.Node, key string) *yaml.Node {
	if node.Kind != yaml.MappingNode {
		return nil
	}
	for i := 0; i+1 < len(node.Content); i += 2 {
		if node.Content[i].Value == key {
			return node.Content[i+1]
		}
	}
	return nil
}

// refs collects the component schema names of every $ref under node.
func refs(node *yaml.Node) []string {
	var out []string
	var walk func(n *yaml.Node)
	walk = func(n *yaml.Node) {
		if n.Kind == yaml.MappingNode {
			for i := 0; i+1 < len(n.Content); i += 2 {
				k, v := n.Content[i], n.Content[i+1]
				if k.Value == "$ref" && v.Kind == yaml.ScalarNode && strings.HasPrefix(v.Value, schemaRefPrefix) {
					out = append(out, strings.TrimPrefix(v.Value, schemaRefPrefix))
					continue
				}
				walk(v)
			}
			return
		}
		for _, c := range n.Content {
			walk(c)
		}
	}
	walk(node)
	return out
}

// typeIndex maps the Go types of the client to the file declaring them.
type typeIndex struct {
	files      map[string]string // Go type → file
	normalised map[string]string // lower-case alphanumerics → Go type
}

func indexTypes(dir string) (*typeIndex, error) {
	idx := &typeIndex{files: map[string]string{}, normalised: map[string]string{}}
	clash := map[string]bool{}
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return err
		}
		file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.SkipObjectResolution)
		if err != nil {
			return fmt.Errorf("%s: parse: %w", path, err)
		}
		for _, decl := range file.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.TYPE {
				continue
			}
			for _, sp := range gen.Specs {
				name := sp.(*ast.TypeSpec).Name.Name
				if !ast.IsExported(name) {
					// _OutputCase, the alias UnmarshalJSON decodes into.
					continue
				}
				idx.files[name] = path
				key := normalise(name)
				if other, taken := idx.normalised[key]; taken && other != name {
					clash[key] = true
				}
				idx.normalised[key] = name
			}
		}
		return nil
	})
	for key := range clash {
		delete(idx.normalised, key)
	}
	return idx, err
}

// lookup finds the Go type openapi-generator emitted for a schema name. It
// camelises names (`_eq` → `Eq`, `case-status` → `CaseStatus`), so the match
// ignores case and anything that is not a letter or a digit.
func (idx *typeIndex) lookup(schema string) (string, bool) {
	if _, ok := idx.files[schema]; ok {
		return schema, true
	}
	goType, ok := idx.normalised[normalise(schema)]
	return goType, ok
}

func normalise(name string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(name) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// wrapper is a naive oneOf wrapper: a struct of pointer variants whose
// UnmarshalJSON counts the variants that decode.
type wrapper struct {
	fields          []field
	namedVariants   []string // the variants that are named model types
	primitiveArrays bool     // every variant is a pointer to a slice of a primitive
	rewritten       bool     // UnmarshalJSON already carries firstMatchMarker
}

type field struct {
	name   string
	goType string // the pointed-to type, e.g. "[]float64"
}

var primitives = map[string]bool{
	"string": true, "bool": true, "int32": true, "int64": true, "float32": true, "float64": true,
}

// inspectWrapper returns nil when goType is not a naive (or already
// first-match) oneOf wrapper.
func inspectWrapper(path, goType string) (*wrapper, error) {
	fset, file, src, err := parseFile(path)
	if err != nil {
		return nil, err
	}
	fd := findUnmarshal(file, goType)
	if fd == nil {
		return nil, nil
	}
	body := rawText(fset, src, fd)
	rewritten := strings.Contains(body, firstMatchMarker)
	naive := strings.Contains(body, "match := 0") &&
		strings.Contains(body, fmt.Sprintf("data matches more than one schema in oneOf(%s)", goType))
	if !rewritten && !naive {
		return nil, nil
	}

	st := findStruct(file, goType)
	if st == nil {
		return nil, fmt.Errorf("oneOf wrapper struct %s not found in %s", goType, path)
	}
	w := &wrapper{primitiveArrays: true, rewritten: rewritten}
	for _, f := range st.Fields.List {
		star, ok := f.Type.(*ast.StarExpr)
		if !ok || len(f.Names) != 1 {
			return nil, fmt.Errorf("%s: field is not a single pointer variant", goType)
		}
		src, err := nodeSource(fset, star.X)
		if err != nil {
			return nil, err
		}
		w.fields = append(w.fields, field{name: f.Names[0].Name, goType: src})
		if ident, ok := star.X.(*ast.Ident); ok && !primitives[ident.Name] {
			w.namedVariants = append(w.namedVariants, ident.Name)
		}
		arr, ok := star.X.(*ast.ArrayType)
		elem, isIdent := (*ast.Ident)(nil), false
		if ok && arr.Len == nil {
			elem, isIdent = arr.Elt.(*ast.Ident)
		}
		if !isIdent || !primitives[elem.Name] {
			w.primitiveArrays = false
		}
	}
	return w, nil
}

// specificity ranks the element type of a primitive-array variant, the most
// specific first: an array of integers also decodes as floats, and an array of
// anything decodes as nothing but strings last.
var specificity = map[string]int{"int32": 0, "int64": 1, "float32": 2, "float64": 3, "bool": 4, "string": 5}

// bySpecificity orders the variants of a primitive-array oneOf, the generator
// having listed them alphabetically.
func bySpecificity(fields []field) []field {
	out := append([]field(nil), fields...)
	rank := func(f field) int { return specificity[strings.TrimPrefix(f.goType, "[]")] }
	sort.SliceStable(out, func(i, j int) bool { return rank(out[i]) < rank(out[j]) })
	return out
}

// The comment lines of a first-match decoder: why counting matches fails there.
var (
	primitiveReason = []string{
		"An empty array, or an array of integers, decodes as several of these",
		"variants at once, and TheHive sends both: counting matches rejects them.",
		"The most specific variant comes first (integers, floats, booleans, strings),",
		"so [1, 2] is an []int32 and not an []float64.",
	}
	namedReason = []string{
		"Each variant checks its own required properties and ignores unknown fields,",
		"so an answer that carries a field TheHive added still matches. The variant",
		"with the most required properties comes first: an answer shaped like it also",
		"satisfies the variants that require less.",
	}
)

// rewriteFirstMatch replaces the naive decoder of a oneOf by one that takes the
// first variant of `ordered` that decodes, and reports whether that changed the
// file. A decoder already rewritten is rewritten again, so a change of order
// reaches the committed client.
func rewriteFirstMatch(path, goType string, ordered []field, reason []string) (bool, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "func (dst *%s) UnmarshalJSON(data []byte) error {\n", goType)
	fmt.Fprintf(&b, "\t// %s.\n", firstMatchMarker)
	for _, line := range reason {
		fmt.Fprintf(&b, "\t// %s\n", line)
	}
	fmt.Fprintf(&b, "\t*dst = %s{}\n", goType)
	for _, f := range ordered {
		v := "v" + f.name
		fmt.Fprintf(&b, "\tvar %s %s\n", v, f.goType)
		fmt.Fprintf(&b, "\tif err := json.Unmarshal(data, &%s); err == nil {\n", v)
		fmt.Fprintf(&b, "\t\tdst.%s = &%s\n\t\treturn nil\n\t}\n", f.name, v)
	}
	fmt.Fprintf(&b, "\treturn fmt.Errorf(\"data failed to match schemas in oneOf(%s)\")\n}\n", goType)

	before, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	if err := spliceFunc(path, goType, b.String(), "gopkg.in/validator.v2"); err != nil {
		return false, err
	}
	after, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	return !bytes.Equal(before, after), nil
}

// shape is what tells a model type apart from the other variants of a naive
// oneOf: what its UnmarshalJSON requires, and what its struct declares.
type shape struct {
	required map[string]bool
	declared map[string]bool
}

// readShape returns nil when goType has no UnmarshalJSON that checks required
// properties, or a field without a JSON name.
func readShape(path, goType string) (*shape, error) {
	_, file, _, err := parseFile(path)
	if err != nil {
		return nil, err
	}
	fd := findUnmarshal(file, goType)
	st := findStruct(file, goType)
	if fd == nil || st == nil {
		return nil, nil
	}

	sh := &shape{required: map[string]bool{}, declared: map[string]bool{}}
	for _, f := range st.Fields.List {
		if f.Tag == nil {
			return nil, nil
		}
		name, _, _ := strings.Cut(reflect.StructTag(strings.Trim(f.Tag.Value, "`")).Get("json"), ",")
		if name == "" {
			return nil, nil
		}
		if name != "-" {
			sh.declared[name] = true
		}
	}

	found := false
	ast.Inspect(fd.Body, func(n ast.Node) bool {
		assign, ok := n.(*ast.AssignStmt)
		if !ok || len(assign.Lhs) != 1 || len(assign.Rhs) != 1 {
			return true
		}
		if id, ok := assign.Lhs[0].(*ast.Ident); !ok || id.Name != "requiredProperties" {
			return true
		}
		lit, ok := assign.Rhs[0].(*ast.CompositeLit)
		if !ok {
			return true
		}
		found = true
		for _, elt := range lit.Elts {
			if bl, ok := elt.(*ast.BasicLit); ok && bl.Kind == token.STRING {
				name, err := strconv.Unquote(bl.Value)
				if err == nil {
					sh.required[name] = true
				}
			}
		}
		return false
	})
	if !found {
		return nil, nil
	}
	return sh, nil
}

// orderVariants returns the variants of a naive oneOf wrapper in the order a
// first-match decoder must try them when they ignore unknown fields, or the
// reason why no order tells them apart. A variant that checks its own required
// properties matches every answer that carries them, so an answer shaped like
// one variant can satisfy another that requires less: the variant with the most
// required properties goes first. It works when no variant tried first has all
// its required properties declared by a variant tried after it, otherwise an
// answer shaped like the later one would be taken for the earlier.
func orderVariants(w *wrapper, files map[string]string) ([]field, string) {
	if len(w.namedVariants) != len(w.fields) {
		return nil, "a variant is not a model type and matches anything the others do"
	}
	shapes := map[string]*shape{}
	for _, f := range w.fields {
		sh, err := readShape(files[f.goType], f.goType)
		if err != nil {
			log.Fatalf("read %s: %v", f.goType, err)
		}
		if sh == nil {
			return nil, fmt.Sprintf("%s does not check required properties", f.goType)
		}
		shapes[f.goType] = sh
	}

	ordered := append([]field(nil), w.fields...)
	sort.SliceStable(ordered, func(i, j int) bool {
		return len(shapes[ordered[i].goType].required) > len(shapes[ordered[j].goType].required)
	})
	for i, first := range ordered {
		for _, later := range ordered[i+1:] {
			covered := true
			for p := range shapes[first.goType].required {
				if !shapes[later.goType].declared[p] {
					covered = false
					break
				}
			}
			if covered {
				return nil, fmt.Sprintf("an answer shaped like %s would be taken for %s", later.goType, first.goType)
			}
		}
	}
	return ordered, ""
}

type result int

const (
	resultNone result = iota
	resultRewritten
	resultAlready
)

// dropDisallow removes the DisallowUnknownFields call from goType's
// UnmarshalJSON. A type without UnmarshalJSON, or whose UnmarshalJSON never
// built a strict decoder, needs nothing.
func dropDisallow(path, goType string) (result, error) {
	fset, file, src, err := parseFile(path)
	if err != nil {
		return resultNone, err
	}
	fd := findUnmarshal(file, goType)
	if fd == nil {
		return resultNone, nil
	}
	body := rawText(fset, src, fd)

	var calls []ast.Stmt
	ast.Inspect(fd.Body, func(n ast.Node) bool {
		stmt, ok := n.(*ast.ExprStmt)
		if !ok {
			return true
		}
		if s, err := nodeSource(fset, stmt); err == nil && s == disallowCall {
			calls = append(calls, stmt)
		}
		return true
	})

	switch {
	case len(calls) == 0 && strings.Contains(body, lenientMarker):
		return resultAlready, nil
	case len(calls) == 0 && strings.Contains(body, "DisallowUnknownFields"):
		return resultNone, fmt.Errorf("strict: UnmarshalJSON in %s disallows unknown fields in a shape this tool does not recognise", path)
	case len(calls) == 0:
		return resultNone, nil
	case len(calls) > 1:
		return resultNone, fmt.Errorf("strict: UnmarshalJSON in %s calls %s %d times, expected once", path, disallowCall, len(calls))
	}

	start := fset.Position(calls[0].Pos()).Offset
	end := fset.Position(calls[0].End()).Offset
	var out bytes.Buffer
	out.Write(src[:start])
	out.WriteString("// " + lenientMarker + ",\n\t// since TheHive adds fields to its responses in minor versions.")
	out.Write(src[end:])
	formatted, err := format.Source(out.Bytes())
	if err != nil {
		return resultNone, fmt.Errorf("format %s: %w", path, err)
	}
	return resultRewritten, os.WriteFile(path, formatted, 0o644)
}

// relaxEnum makes the generated UnmarshalJSON of a string enum keep a value it
// does not list instead of failing the whole answer. The template accepts only
// the values of the spec the client was generated from; TheHive adds values to
// its enums in minor versions. A type that is not such an enum needs nothing.
func relaxEnum(path, goType string) (result, error) {
	fset, file, src, err := parseFile(path)
	if err != nil {
		return resultNone, err
	}
	fd := findUnmarshal(file, goType)
	if fd == nil {
		return resultNone, nil
	}
	body := rawText(fset, src, fd)

	switch {
	case strings.Contains(body, enumMarker):
		return resultAlready, nil
	case !strings.Contains(body, enumRejection(goType)):
		return resultNone, nil
	case !strings.Contains(body, "var value string") || !strings.Contains(body, fmt.Sprintf("enumTypeValue := %s(value)", goType)):
		return resultNone, fmt.Errorf("strict: the enum UnmarshalJSON in %s has a shape this tool does not recognise (it rewrites string enums)", path)
	}

	var b strings.Builder
	fmt.Fprintf(&b, "func (v *%s) UnmarshalJSON(src []byte) error {\n", goType)
	fmt.Fprintf(&b, "\t// %s,\n", enumMarker)
	b.WriteString("\t// since TheHive adds values to its enums in minor versions.\n")
	b.WriteString("\tvar value string\n")
	b.WriteString("\tif err := json.Unmarshal(src, &value); err != nil {\n\t\treturn err\n\t}\n")
	fmt.Fprintf(&b, "\t*v = %s(value)\n\treturn nil\n}\n", goType)
	return resultRewritten, spliceFunc(path, goType, b.String(), "gopkg.in/validator.v2")
}

// enumRejection is the end of the error a generated enum decoder returns for a
// value it does not list.
func enumRejection(goType string) string {
	return fmt.Sprintf("is not a valid %s\"", goType)
}

// verify fails when a response model still disallows unknown fields without
// being a variant of a naive oneOf, when a response enum still rejects a value
// it does not list, or when a discriminated oneOf dispatches to a model without
// an UnmarshalJSON of its own: newStrictDecoder would decode it with
// DisallowUnknownFields, and no rewrite above could reach it.
func verify(responseTypes []string, strictVariants map[string]string, files map[string]string) error {
	var left, orphans, enums []string
	for _, goType := range responseTypes {
		fset, file, src, err := parseFile(files[goType])
		if err != nil {
			return err
		}
		fd := findUnmarshal(file, goType)
		if fd == nil {
			continue
		}
		body := rawText(fset, src, fd)

		if strings.Contains(body, dispatchMarker) {
			for _, target := range dispatchTargets(fd) {
				path, ok := files[target]
				if !ok {
					return fmt.Errorf("%s dispatches to %s, which no Go type declares", goType, target)
				}
				_, targetFile, _, err := parseFile(path)
				if err != nil {
					return err
				}
				if findUnmarshal(targetFile, target) == nil {
					orphans = append(orphans, goType+" → "+target)
				}
			}
		}

		if _, kept := strictVariants[goType]; !kept && strings.Contains(body, "DisallowUnknownFields()") {
			left = append(left, goType)
		}
		if strings.Contains(body, enumRejection(goType)) {
			enums = append(enums, goType)
		}
	}
	if len(enums) > 0 {
		return fmt.Errorf("response enum(s) still reject a value they do not list: %s", strings.Join(enums, ", "))
	}
	if len(orphans) > 0 {
		return fmt.Errorf("discriminated oneOf dispatch target(s) without an UnmarshalJSON of their own, which newStrictDecoder decodes strictly: %s",
			strings.Join(orphans, ", "))
	}
	if len(left) > 0 {
		return fmt.Errorf("response model(s) still disallow unknown fields: %s", strings.Join(left, ", "))
	}
	return nil
}

// dispatchTargets lists the model types a discriminated oneOf decoder
// allocates: `dst.Field = new(Type)`.
func dispatchTargets(fd *ast.FuncDecl) []string {
	var out []string
	ast.Inspect(fd.Body, func(n ast.Node) bool {
		assign, ok := n.(*ast.AssignStmt)
		if !ok || len(assign.Rhs) != 1 {
			return true
		}
		call, ok := assign.Rhs[0].(*ast.CallExpr)
		if !ok || len(call.Args) != 1 {
			return true
		}
		if fn, ok := call.Fun.(*ast.Ident); !ok || fn.Name != "new" {
			return true
		}
		if arg, ok := call.Args[0].(*ast.Ident); ok {
			out = append(out, arg.Name)
		}
		return true
	})
	return out
}

func findUnmarshal(file *ast.File, goType string) *ast.FuncDecl {
	for _, decl := range file.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok || fd.Name.Name != "UnmarshalJSON" || fd.Recv == nil || len(fd.Recv.List) != 1 {
			continue
		}
		star, ok := fd.Recv.List[0].Type.(*ast.StarExpr)
		if !ok {
			continue
		}
		if ident, ok := star.X.(*ast.Ident); ok && ident.Name == goType {
			return fd
		}
	}
	return nil
}

func findStruct(file *ast.File, goType string) *ast.StructType {
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.TYPE {
			continue
		}
		for _, sp := range gen.Specs {
			ts := sp.(*ast.TypeSpec)
			if ts.Name.Name != goType {
				continue
			}
			if st, ok := ts.Type.(*ast.StructType); ok && st.Fields != nil {
				return st
			}
		}
	}
	return nil
}

func parseFile(path string) (*token.FileSet, *ast.File, []byte, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, nil, err
	}
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, src, parser.ParseComments)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("parse %s: %w", path, err)
	}
	return fset, file, src, nil
}

// rawText is node's source as written, comments included (format.Node drops
// the comments inside a function body, and the markers are comments).
func rawText(fset *token.FileSet, src []byte, node ast.Node) string {
	return string(src[fset.Position(node.Pos()).Offset:fset.Position(node.End()).Offset])
}

func nodeSource(fset *token.FileSet, node ast.Node) (string, error) {
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, node); err != nil {
		return "", fmt.Errorf("format node: %w", err)
	}
	return buf.String(), nil
}

// spliceFunc replaces goType's UnmarshalJSON (and its doc comment) with
// newSource, drops an import the new code no longer uses, and gofmts.
func spliceFunc(path, goType, newSource, unusedImport string) error {
	fset := token.NewFileSet()
	src, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	file, err := parser.ParseFile(fset, path, src, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}
	fd := findUnmarshal(file, goType)
	if fd == nil {
		return fmt.Errorf("UnmarshalJSON for %s not found", goType)
	}
	var spliced bytes.Buffer
	spliced.Write(src[:fset.Position(fd.Pos()).Offset])
	spliced.WriteString(newSource)
	spliced.Write(src[fset.Position(fd.End()).Offset:])

	fset2 := token.NewFileSet()
	file2, err := parser.ParseFile(fset2, path, spliced.Bytes(), parser.ParseComments)
	if err != nil {
		return fmt.Errorf("re-parse: %w", err)
	}
	if !usesPackage(file2, "validator") {
		removeImport(file2, unusedImport)
	}
	var out bytes.Buffer
	if err := format.Node(&out, fset2, file2); err != nil {
		return fmt.Errorf("format: %w", err)
	}
	return os.WriteFile(path, out.Bytes(), 0o644)
}

// usesPackage reports whether the file refers to pkg.Something anywhere.
func usesPackage(file *ast.File, pkg string) bool {
	used := false
	ast.Inspect(file, func(n ast.Node) bool {
		if sel, ok := n.(*ast.SelectorExpr); ok {
			if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == pkg {
				used = true
			}
		}
		return !used
	})
	return used
}

func removeImport(file *ast.File, importPath string) {
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.IMPORT {
			continue
		}
		kept := gen.Specs[:0]
		for _, sp := range gen.Specs {
			if is, ok := sp.(*ast.ImportSpec); ok && is.Path.Value == `"`+importPath+`"` {
				continue
			}
			kept = append(kept, sp)
		}
		gen.Specs = kept
	}
}
