package regression

import (
	"encoding/json"
	"testing"

	"github.com/StrangeBeeCorp/thehive4go/thehive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestResponseEnumsKeepAValueTheyDoNotList pins scripts/tolerant-response-decoder
// on enums. The generated UnmarshalJSON of an enum accepts only the values of the
// spec the client was generated from and fails the whole answer on another one,
// which a TheHive that adds a permission, a custom field type or a profile type
// in a minor version would trigger.
func TestResponseEnumsKeepAValueTheyDoNotList(t *testing.T) {
	t.Run("an enum only responses use: PermissionDesc.Name", func(t *testing.T) {
		var p thehive.PermissionDesc
		require.NoError(t, json.Unmarshal([]byte(`{"name":"manageSomethingNew","label":"Manage something new","consumesLicense":false}`), &p))

		assert.Equal(t, "manageSomethingNew", string(p.Name))
	})

	t.Run("an enum a request shares: OutputCustomField.Type", func(t *testing.T) {
		payload := `{"_id":"~1","_type":"CustomField","_createdBy":"admin@thehive.local","_createdAt":1,
			"name":"f","displayName":"F","group":"default","description":"d","type":"quantum",
			"options":[],"mandatory":false,"extraData":{}}`
		var field thehive.OutputCustomField
		require.NoError(t, json.Unmarshal([]byte(payload), &field))

		assert.Equal(t, "quantum", string(field.Type))
	})

	t.Run("a value the enum lists still decodes", func(t *testing.T) {
		var p thehive.PermissionDesc
		require.NoError(t, json.Unmarshal([]byte(`{"name":"manageCase/create","label":"Create cases","consumesLicense":false}`), &p))

		assert.Equal(t, thehive.Permissions("manageCase/create"), p.Name)
	})
}

// TestRequestOnlyEnumsStayStrict checks that the enums only requests refer to
// still reject a value they do not list: a consumer that decodes a caller's input
// into the request models relies on it.
func TestRequestOnlyEnumsStayStrict(t *testing.T) {
	var stage thehive.InputCaseStage
	err := json.Unmarshal([]byte(`"NotAStage"`), &stage)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "is not a valid InputCaseStage")
}
