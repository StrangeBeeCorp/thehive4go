// Package regression holds hand-maintained tests that pin known bugs in the
// generated SDK. These tests live outside thehive/ on purpose: that directory
// is wiped on every `make generate` (see scripts/generate_client.sh), so any
// hand-written test placed there would be lost. The tests here survive
// regeneration and act as a guard against the generator (or a postprocess
// step) silently regressing.
//
// Today this package contains a single regression: the Access oneOf decoder
// for TheHive 5.6.
package regression

import (
	"encoding/json"
	"testing"

	"github.com/StrangeBeeCorp/thehive4go/thehive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAccessOneOfDecoder pins the bug introduced when TheHive 5.6 added the
// AllExternalAccess and ExternalAccess variants to the Access oneOf.
//
// The four variants split into two pairs of byte-identical Go structs:
//
//	OrganisationAccess ≡ AllExternalAccess  →  { _kind string }
//	UserAccess         ≡ ExternalAccess     →  { _kind string, users []string }
//
// The generated Access.UnmarshalJSON does naive structural matching against
// every variant and counts matches, ignoring the OpenAPI discriminator on
// _kind. Within each pair both variants decode successfully, match counter
// reaches 2, and the decoder returns: "data matches more than one schema in
// oneOf(Access)". Every endpoint returning OutputCase is broken as a result.
//
// Expected once a fix lands (likely a postprocess patch that honors the
// _kind discriminator): each payload resolves to exactly one variant.
func TestAccessOneOfDecoder(t *testing.T) {
	t.Run("OrganisationAccess payload resolves to OrganisationAccess only", func(t *testing.T) {
		var access thehive.Access
		err := json.Unmarshal([]byte(`{"_kind":"OrganisationAccessKind"}`), &access)

		require.NoError(t, err)
		require.NotNil(t, access.OrganisationAccess)
		assert.Equal(t, "OrganisationAccessKind", access.OrganisationAccess.GetKind())
		assert.Nil(t, access.AllExternalAccess)
		assert.Nil(t, access.UserAccess)
		assert.Nil(t, access.ExternalAccess)
	})

	t.Run("AllExternalAccess payload resolves to AllExternalAccess only", func(t *testing.T) {
		var access thehive.Access
		err := json.Unmarshal([]byte(`{"_kind":"AllExternalAccessKind"}`), &access)

		require.NoError(t, err)
		require.NotNil(t, access.AllExternalAccess)
		assert.Equal(t, "AllExternalAccessKind", access.AllExternalAccess.GetKind())
		assert.Nil(t, access.OrganisationAccess)
		assert.Nil(t, access.UserAccess)
		assert.Nil(t, access.ExternalAccess)
	})

	t.Run("UserAccess payload resolves to UserAccess only", func(t *testing.T) {
		var access thehive.Access
		err := json.Unmarshal([]byte(`{"_kind":"UserAccessKind","users":["alice","bob"]}`), &access)

		require.NoError(t, err)
		require.NotNil(t, access.UserAccess)
		assert.Equal(t, "UserAccessKind", access.UserAccess.GetKind())
		assert.Equal(t, []string{"alice", "bob"}, access.UserAccess.GetUsers())
		assert.Nil(t, access.OrganisationAccess)
		assert.Nil(t, access.AllExternalAccess)
		assert.Nil(t, access.ExternalAccess)
	})

	t.Run("ExternalAccess payload resolves to ExternalAccess only", func(t *testing.T) {
		var access thehive.Access
		err := json.Unmarshal([]byte(`{"_kind":"ExternalAccessKind","users":["carol"]}`), &access)

		require.NoError(t, err)
		require.NotNil(t, access.ExternalAccess)
		assert.Equal(t, "ExternalAccessKind", access.ExternalAccess.GetKind())
		assert.Equal(t, []string{"carol"}, access.ExternalAccess.GetUsers())
		assert.Nil(t, access.OrganisationAccess)
		assert.Nil(t, access.AllExternalAccess)
		assert.Nil(t, access.UserAccess)
	})
}

// TestOutputCaseDecodesAccessField ensures the bug is fixed at the level
// where it actually bites users: every OutputCase returned by TheHive 5.6
// embeds an `access` object, so any endpoint returning OutputCase (CreateCase,
// GetCase, listCase queries, ...) hits this decoder. Payload mirrors what a
// stock TheHive 5.6 instance returns for a freshly created case in a default
// organisation.
func TestOutputCaseDecodesAccessField(t *testing.T) {
	payload := []byte(`{
		"_id":"~123",
		"_type":"Case",
		"_createdBy":"admin@thehive.local",
		"_createdAt":1700000000000,
		"number":1,
		"title":"regression",
		"description":"regression",
		"severity":2,
		"severityLabel":"MEDIUM",
		"startDate":1700000000000,
		"flag":false,
		"tlp":2,
		"tlpLabel":"AMBER",
		"pap":2,
		"papLabel":"AMBER",
		"status":"New",
		"stage":"New",
		"access":{"_kind":"OrganisationAccessKind"},
		"extraData":{},
		"newDate":1700000000000,
		"timeToDetect":0
	}`)

	var c thehive.OutputCase
	err := json.Unmarshal(payload, &c)

	require.NoError(t, err)
	access := c.GetAccess()
	require.NotNil(t, access.OrganisationAccess)
	assert.Equal(t, "OrganisationAccessKind", access.OrganisationAccess.GetKind())
}
