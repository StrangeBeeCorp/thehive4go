package integration

import (
	"fmt"
	"testing"
	"time"

	"github.com/StrangeBeeCorp/thehive4go/tests/integration/testutils"
	"github.com/StrangeBeeCorp/thehive4go/thehive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestListCustomFieldsDecodesEveryOptionShape lists the custom fields once the
// instance holds one without options and one with integer options. Before
// scripts/tolerant-response-decoder, both made the whole list fail with
// "data matches more than one schema in oneOf(OutputCustomFieldOptions)".
func TestListCustomFieldsDecodesEveryOptionShape(t *testing.T) {
	cfg := testutils.GetTestConfig(t)
	client := testutils.CreateAdminClient(t, cfg)
	ctx := testutils.CreateAuthContext(t, cfg)

	suffix := time.Now().UnixNano()
	fields := []struct {
		name    string
		type_   thehive.CustomFieldType
		options []interface{}
	}{
		{fmt.Sprintf("t4g-no-options-%d", suffix), thehive.CUSTOMFIELDTYPE_STRING, nil},
		{fmt.Sprintf("t4g-integer-options-%d", suffix), thehive.CUSTOMFIELDTYPE_INTEGER, []interface{}{1, 2, 3}},
	}

	for _, f := range fields {
		input := thehive.NewInputCustomField(f.name, "thehive4go", "integration test", f.type_)
		if f.options != nil {
			input.SetOptions(f.options)
		}
		created, _, err := client.CustomFieldAPI.CreateCustomField(ctx).InputCustomField(*input).Execute()
		require.NoError(t, err, "create custom field %s", f.name)
		t.Cleanup(func() {
			_, _ = client.CustomFieldAPI.DeleteCustomField(ctx, created.GetUnderscoreId()).Execute()
		})
	}

	listed, _, err := client.CustomFieldAPI.ListCustomFields(ctx).Execute()
	require.NoError(t, err)

	byName := map[string]thehive.OutputCustomField{}
	for _, f := range listed {
		byName[f.GetName()] = f
	}
	for _, f := range fields {
		got, ok := byName[f.name]
		require.True(t, ok, "custom field %s not listed", f.name)
		want := "[]"
		if f.options != nil {
			want = "[1,2,3]"
		}
		out, err := got.Options.MarshalJSON()
		require.NoError(t, err)
		assert.JSONEq(t, want, string(out))
	}
}
