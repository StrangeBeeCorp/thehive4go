package regression

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/StrangeBeeCorp/thehive4go/thehive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestResponseModelsIgnoreUnknownFields pins the fields TheHive added to its
// responses after 5.6, the spec the client is generated from. The generator
// decodes models with DisallowUnknownFields, so before
// scripts/tolerant-response-decoder every TheHive 5.8 responder action failed
// with `json: unknown field "extraData"`, although TheHive had created it.
func TestResponseModelsIgnoreUnknownFields(t *testing.T) {
	cases := []struct {
		name    string
		payload string
		decode  func([]byte) error
	}{
		{
			// A responder action as TheHive 6.0.0-SNAPSHOT answers it; extraData since 5.8.
			name: "OutputAction with extraData (5.8)",
			payload: `{"_createdAt":1791986098031,"_createdBy":"analyst@example.org","_id":"~2416713896","_type":"Action",
				"cortexId":"local","cortexJobId":"0lmxOqEBQp5xTN6g36nT","endDate":1791986103671,"extraData":{},
				"objectId":"~2703511728","objectType":"Task",
				"operations":"[{\"type\":\"CloseTask\",\"status\":\"Success\",\"message\":\"Success\"}]",
				"report":"{\"summary\":{\"taxonomies\":[]},\"full\":{\"message\":\"OK\"},\"success\":true}",
				"responderDefinition":"testResponder_1_1_0","responderId":"bdfa784dc600c39e27ca89187a4a1662",
				"responderName":"testResponder_1_1_0","startDate":1791986098130,"status":"Success"}`,
			decode: func(b []byte) error { var v thehive.OutputAction; return json.Unmarshal(b, &v) },
		},
		{
			name: "OutputTask with timeToHandle (5.8)",
			payload: `{"_id":"~1","_type":"Task","_createdBy":"analyst@example.org","_createdAt":1,"title":"Triage",
				"group":"default","status":"Completed","flag":false,"order":0,"mandatory":false,"extraData":{},
				"timeToHandle":42}`,
			decode: func(b []byte) error { var v thehive.OutputTask; return json.Unmarshal(b, &v) },
		},
		{
			name: "OutputTag with hidden (5.8)",
			payload: `{"_id":"~1","_type":"Tag","_createdBy":"analyst@example.org","_createdAt":1,"namespace":"ns",
				"predicate":"p","colour":"#000000","extraData":{},"hidden":false}`,
			decode: func(b []byte) error { var v thehive.OutputTag; return json.Unmarshal(b, &v) },
		},
		{
			name: "OutputObservableType with isCaseSensitive (5.8)",
			payload: `{"_id":"~1","_type":"ObservableType","_createdBy":"admin@thehive.local","_createdAt":1,
				"name":"domain","isAttachment":false,"isCaseSensitive":false}`,
			decode: func(b []byte) error { var v thehive.OutputObservableType; return json.Unmarshal(b, &v) },
		},
		{
			name: "OutputCaseStatus with requiresImpact (6.0)",
			payload: `{"_id":"~1","_type":"CaseStatus","_createdBy":"admin@thehive.local","_createdAt":1,
				"value":"TruePositive","stage":"Closed","hidden":false,"extraData":{},"requiresImpact":true}`,
			decode: func(b []byte) error { var v thehive.OutputCaseStatus; return json.Unmarshal(b, &v) },
		},
		{
			// A nested model, reached through the Access oneOf of OutputCase.
			name:    "Access variant with an unknown field",
			payload: `{"_kind":"OrganisationAccessKind","addedIn":"6.0"}`,
			decode:  func(b []byte) error { var v thehive.Access; return json.Unmarshal(b, &v) },
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.NoError(t, tc.decode([]byte(tc.payload)))
		})
	}
}

// TestResponseModelsStillRequireTheirFields checks that tolerating unknown
// fields did not also drop the required-property check.
func TestResponseModelsStillRequireTheirFields(t *testing.T) {
	var action thehive.OutputAction
	err := json.Unmarshal([]byte(`{"_id":"~1","extraData":{}}`), &action)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "no value given for required property")
}

// TestOlderAnswersDecodeWithoutTheFieldsTheyLack pins the required fields that
// scripts/preprocess_openapi.sh relaxes because only newer TheHive versions send
// them: extraData on actions, hidden on tags and isCaseSensitive on observable
// types (5.8), requiresImpact on case statuses (6.0). The 6.0 spec lists them as
// required; without the relaxation every answer of an older TheHive fails with
// "no value given for required property".
func TestOlderAnswersDecodeWithoutTheFieldsTheyLack(t *testing.T) {
	cases := []struct {
		name    string
		payload string
		// absent decodes the payload and reports whether the newer field stayed unset.
		absent func([]byte) (bool, error)
	}{
		{
			name: "OutputAction without extraData (before 5.8)",
			payload: `{"_createdAt":1791986098031,"_createdBy":"analyst@example.org","_id":"~2416713896","_type":"Action",
				"cortexId":"local","cortexJobId":"0lmxOqEBQp5xTN6g36nT","endDate":1791986103671,
				"objectId":"~2703511728","objectType":"Task","operations":"[]","report":"{}",
				"responderDefinition":"testResponder_1_1_0","responderId":"bdfa784dc600c39e27ca89187a4a1662",
				"responderName":"testResponder_1_1_0","startDate":1791986098130,"status":"Success"}`,
			absent: func(b []byte) (bool, error) {
				var v thehive.OutputAction
				err := json.Unmarshal(b, &v)
				return v.ExtraData == nil, err
			},
		},
		{
			name: "OutputTag without hidden (before 5.8)",
			payload: `{"_id":"~1","_type":"Tag","_createdBy":"analyst@example.org","_createdAt":1,"namespace":"ns",
				"predicate":"p","colour":"#000000","extraData":{}}`,
			absent: func(b []byte) (bool, error) {
				var v thehive.OutputTag
				err := json.Unmarshal(b, &v)
				return v.Hidden == nil, err
			},
		},
		{
			name: "OutputObservableType without isCaseSensitive (before 5.8)",
			payload: `{"_id":"~1","_type":"ObservableType","_createdBy":"admin@thehive.local","_createdAt":1,
				"name":"domain","isAttachment":false}`,
			absent: func(b []byte) (bool, error) {
				var v thehive.OutputObservableType
				err := json.Unmarshal(b, &v)
				return v.IsCaseSensitive == nil, err
			},
		},
		{
			name: "OutputCaseStatus without requiresImpact (before 6.0)",
			payload: `{"_id":"~1","_type":"CaseStatus","_createdBy":"admin@thehive.local","_createdAt":1,
				"value":"TruePositive","stage":"Closed","hidden":false,"extraData":{}}`,
			absent: func(b []byte) (bool, error) {
				var v thehive.OutputCaseStatus
				err := json.Unmarshal(b, &v)
				return v.RequiresImpact == nil, err
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			absent, err := tc.absent([]byte(tc.payload))
			require.NoError(t, err)
			assert.True(t, absent, "the field the older answer lacks must stay unset")
		})
	}
}

// customFieldOptionsVariant names the one variant of the oneOf that is set.
func customFieldOptionsVariant(t *testing.T, o thehive.OutputCustomFieldOptions) string {
	t.Helper()
	var set []string
	if o.ArrayOfInt32 != nil {
		set = append(set, "ArrayOfInt32")
	}
	if o.ArrayOfFloat64 != nil {
		set = append(set, "ArrayOfFloat64")
	}
	if o.ArrayOfString != nil {
		set = append(set, "ArrayOfString")
	}
	require.Len(t, set, 1, "exactly one variant must be set")
	return set[0]
}

// TestCustomFieldOptionsDecodeEveryShapeTheHiveSends pins the naive oneOf
// decoder of OutputCustomFieldOptions ([]float64 | []int32 | []string). It
// counted the variants that decode and failed on more than one, so `[]` (all
// three) and integer options (float and int) broke
// GET /api/v1/customField for most organisations. The first variant that
// decodes wins, the most specific first: integer options must not come out as
// floats.
func TestCustomFieldOptionsDecodeEveryShapeTheHiveSends(t *testing.T) {
	cases := []struct {
		name    string
		fieldTy string
		options string
		variant string
	}{
		{"no options", "string", `[]`, "ArrayOfInt32"},
		{"null options", "string", `null`, "ArrayOfInt32"},
		{"integer options", "integer", `[0,1,2,3,10]`, "ArrayOfInt32"},
		{"float options", "float", `[1.5,2]`, "ArrayOfFloat64"},
		// Past int32: the integers fall to the float variant, and survive it.
		{"integer options past int32", "integer", `[3000000000]`, "ArrayOfFloat64"},
		{"string options", "string", `["HR","Sales"]`, "ArrayOfString"},
		{"one string option", "string", `["a"]`, "ArrayOfString"},
		// TheHive keeps the options when a field's type changes.
		{"string options on an integer field", "integer", `["Malware","Intrusion"]`, "ArrayOfString"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			payload := `{"_id":"~1","_type":"CustomField","_createdBy":"admin@thehive.local","_createdAt":1,
				"name":"f","displayName":"F","group":"default","description":"d","type":"` + tc.fieldTy + `",
				"options":` + tc.options + `,"mandatory":false,"extraData":{}}`

			var field thehive.OutputCustomField
			require.NoError(t, json.Unmarshal([]byte(payload), &field))

			assert.Equal(t, tc.variant, customFieldOptionsVariant(t, field.Options))
			out, err := json.Marshal(field.Options)
			require.NoError(t, err)
			assert.JSONEq(t, tc.options, string(out), "options must survive a decode and an encode unchanged")
		})
	}

	t.Run("mixed options match no variant", func(t *testing.T) {
		var options thehive.OutputCustomFieldOptions
		assert.Error(t, json.Unmarshal([]byte(`[1,"a"]`), &options))
	})
}

// TestEnumerationValuesDecodeEveryShapeTheHiveSends pins the same fix on the
// other primitive-array oneOf, the values of an enumeration property
// description ([]int32 | []string).
func TestEnumerationValuesDecodeEveryShapeTheHiveSends(t *testing.T) {
	cases := []struct {
		name   string
		values string
		isInt  bool
	}{
		{"no values", `[]`, true},
		{"integer values", `[1,2]`, true},
		{"string values", `["open","closed"]`, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var values thehive.EnumerationPropertyDescriptionValues
			require.NoError(t, json.Unmarshal([]byte(tc.values), &values))

			assert.Equal(t, tc.isInt, values.ArrayOfInt32 != nil)
			assert.Equal(t, !tc.isInt, values.ArrayOfString != nil)
			out, err := json.Marshal(values)
			require.NoError(t, err)
			assert.JSONEq(t, tc.values, string(out))
		})
	}
}

// TestLicenseCurrentDecodesWithAnUnknownField pins the variants of
// OutputLicenseCurrent, a naive oneOf no request refers to. Each variant
// requires a property the others do not declare, so they stay apart without
// DisallowUnknownFields; before, `{"license":{...},"newField":1}` failed
// GetCurrentLicense.
func TestLicenseCurrentDecodesWithAnUnknownField(t *testing.T) {
	const license = `{"_id":"~1","_type":"License","_createdBy":"admin@thehive.local","_createdAt":1,"id":"l1",
		"customer":"c","plan":"p","kind":"k","validFrom":1,"expiresAt":2,"current":true,"quotas":{"users":10}}`

	cases := []struct {
		name    string
		payload string
		variant string
	}{
		{"license", `"license":` + license, "Ok"},
		{"error with a fallback", `"error":"unreadable","fallback":` + license, "Error"},
		{"no license found", `"notFound":true,"fallback":` + license, "NotFound"},
	}

	for _, tc := range cases {
		for _, extra := range []string{"", `,"newField":1`} {
			t.Run(tc.name+extra, func(t *testing.T) {
				var current thehive.OutputLicenseCurrent
				require.NoError(t, json.Unmarshal([]byte(`{`+tc.payload+extra+`}`), &current))

				assert.Equal(t, tc.variant == "Ok", current.OutputLicenseCurrentOk != nil)
				assert.Equal(t, tc.variant == "Error", current.OutputLicenseCurrentError != nil)
				assert.Equal(t, tc.variant == "NotFound", current.OutputLicenseCurrentNotFound != nil)
			})
		}
	}
}

// TestNoResponseModelDisallowsUnknownFields guards the generated sources: a
// regeneration that skips scripts/tolerant-response-decoder brings
// DisallowUnknownFields back into the Output models. The variants of
// OutputInvokeFunction keep it: its map variant matches any object, so a
// lenient Ok or Ko would make an answer with an unknown field match two
// variants and fail.
func TestNoResponseModelDisallowsUnknownFields(t *testing.T) {
	keptStrict := map[string]bool{
		"model_output_invoke_function_ko.go": true,
		"model_output_invoke_function_ok.go": true,
	}

	files, err := filepath.Glob(filepath.Join("..", "..", "thehive", "model_output_*.go"))
	require.NoError(t, err)
	require.NotEmpty(t, files, "the generated Output models were not found")

	strict := regexp.MustCompile(`(?m)^\s*decoder\.DisallowUnknownFields\(\)`)
	var offenders []string
	for _, f := range files {
		if keptStrict[filepath.Base(f)] {
			continue
		}
		src, err := os.ReadFile(f)
		require.NoError(t, err)
		if strict.Match(src) {
			offenders = append(offenders, filepath.Base(f))
		}
	}

	assert.Empty(t, offenders, "these response models reject unknown fields again: run make generate, whose "+
		"scripts/tolerant-response-decoder step makes them lenient:\n%s", strings.Join(offenders, "\n"))
}
