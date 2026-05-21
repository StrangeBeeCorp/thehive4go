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

// TestOneOfWrapperSmoke covers the five other wrappers the fix-oneof-decoder
// tool rewrites alongside Access (Auth, InputEmailIntakeMailboxConfig,
// OutputEmailIntakeMailbox, PropertyDescription, Widget). For each one we
// feed a minimal payload carrying just the discriminator, confirm the
// dispatcher routes to the expected variant field and leaves the others
// nil. If the AST rewriter ever miswires a wrapper — wrong field, wrong
// discriminator key, missing case — these tests catch it before CI green-
// lights a broken release.
func TestOneOfWrapperSmoke(t *testing.T) {
	t.Run("Auth dispatches basic to BasicAuthCredentials", func(t *testing.T) {
		var a thehive.Auth
		require.NoError(t, json.Unmarshal([]byte(`{"type":"basic","username":"u","password":"p"}`), &a))
		require.NotNil(t, a.BasicAuthCredentials)
		assert.Equal(t, "basic", a.BasicAuthCredentials.Type)
		assert.Equal(t, "u", a.BasicAuthCredentials.Username)
		assert.Nil(t, a.BearerAuth)
		assert.Nil(t, a.KeyAuth)
		assert.Nil(t, a.NoneAuth)
	})

	t.Run("InputEmailIntakeMailboxConfig dispatches api to InputEmailIntakeApiMailbox", func(t *testing.T) {
		payload := []byte(`{
			"_kind": "api",
			"provider": {"name": "office365"},
			"credential": {"email": "test@example.com"}
		}`)
		var c thehive.InputEmailIntakeMailboxConfig
		require.NoError(t, json.Unmarshal(payload, &c))
		require.NotNil(t, c.InputEmailIntakeApiMailbox)
		require.NotNil(t, c.InputEmailIntakeApiMailbox.Kind)
		assert.Equal(t, "api", *c.InputEmailIntakeApiMailbox.Kind)
		assert.Nil(t, c.InputEmailIntakeImapMailbox)
	})

	t.Run("OutputEmailIntakeMailbox dispatches imap to OutputEmailIntakeImapMailbox", func(t *testing.T) {
		payload := []byte(`{
			"_kind": "imap",
			"provider": {"name": "office365"},
			"credential": {"email": "test@example.com"},
			"inbox": "INBOX",
			"markAsRead": true
		}`)
		var m thehive.OutputEmailIntakeMailbox
		require.NoError(t, json.Unmarshal(payload, &m))
		require.NotNil(t, m.OutputEmailIntakeImapMailbox)
		assert.Equal(t, "imap", m.OutputEmailIntakeImapMailbox.Kind)
		assert.Nil(t, m.OutputEmailIntakeApiMailbox)
	})

	t.Run("PropertyDescription dispatches boolean to BooleanPropertyDescription", func(t *testing.T) {
		payload := []byte(`{
			"type": "boolean",
			"name": "isActive",
			"cardinality": "single",
			"aggregable": false,
			"indexType": "standard"
		}`)
		var p thehive.PropertyDescription
		require.NoError(t, json.Unmarshal(payload, &p))
		require.NotNil(t, p.BooleanPropertyDescription)
		assert.Equal(t, "boolean", p.BooleanPropertyDescription.Type)
		assert.Equal(t, "isActive", p.BooleanPropertyDescription.Name)
		assert.Nil(t, p.DatePropertyDescription)
		assert.Nil(t, p.EnumerationPropertyDescription)
		assert.Nil(t, p.FloatPropertyDescription)
		assert.Nil(t, p.IntegerPropertyDescription)
		assert.Nil(t, p.StringPropertyDescription)
		assert.Nil(t, p.UrlPropertyDescription)
		assert.Nil(t, p.UserPropertyDescription)
	})

	t.Run("Widget dispatches Comments to Comments", func(t *testing.T) {
		var w thehive.Widget
		require.NoError(t, json.Unmarshal([]byte(`{"_kind":"Comments"}`), &w))
		require.NotNil(t, w.Comments)
		assert.Equal(t, thehive.Kind("Comments"), w.Comments.Kind)
		assert.Nil(t, w.AlertList)
		assert.Nil(t, w.AlertTable)
		assert.Nil(t, w.Image)
		assert.Nil(t, w.Text)
		assert.Nil(t, w.Timeline)
	})
}

// TestRelaxedRequiredFields55Compatibility pins the cross-version 5.5/5.6
// SDK contract: scripts/preprocess_openapi.sh drops five fields from the
// `required:` list of as many models so that 5.5 payloads (which don't
// carry them) decode successfully. If a future regeneration silently
// re-introduces any of these as required — because the preprocess awk
// stopped matching, or the spec changed shape — the matching subtest
// here fires immediately, before a 5.5 deployment hits a "no value given
// for required property X" decode failure in prod.
//
// Each subtest feeds the minimal payload TheHive 5.5 actually emits for
// that model (no relaxed field present) and asserts the decode succeeds
// AND the relaxed field is left nil.
func TestRelaxedRequiredFields55Compatibility(t *testing.T) {
	t.Run("OutputComment decodes without external", func(t *testing.T) {
		payload := []byte(`{
			"_id": "~1",
			"_type": "Comment",
			"createdBy": "admin@thehive.local",
			"createdAt": 1700000000000,
			"message": "hello",
			"isEdited": false,
			"extraData": {}
		}`)
		var c thehive.OutputComment
		require.NoError(t, json.Unmarshal(payload, &c))
		assert.Equal(t, "hello", c.Message)
		assert.Nil(t, c.External, "External must be optional for 5.5 compatibility")
	})

	t.Run("OutputObservable decodes without external", func(t *testing.T) {
		payload := []byte(`{
			"_id": "~1",
			"_type": "Observable",
			"_createdBy": "admin@thehive.local",
			"_createdAt": 1700000000000,
			"dataType": "ip",
			"startDate": 1700000000000,
			"tlp": 2,
			"tlpLabel": "AMBER",
			"pap": 2,
			"papLabel": "AMBER",
			"ioc": false,
			"sighted": false,
			"reports": {},
			"extraData": {},
			"ignoreSimilarity": false
		}`)
		var o thehive.OutputObservable
		require.NoError(t, json.Unmarshal(payload, &o))
		assert.Equal(t, "ip", o.DataType)
		assert.Nil(t, o.External, "External must be optional for 5.5 compatibility")
	})

	t.Run("OutputAttachment decodes without external", func(t *testing.T) {
		payload := []byte(`{
			"_id": "~1",
			"_type": "Attachment",
			"_createdBy": "admin@thehive.local",
			"_createdAt": 1700000000000,
			"name": "f.txt",
			"hashes": ["sha256:abc"],
			"size": 1,
			"contentType": "text/plain",
			"id": "id-1",
			"path": "/tmp/f.txt",
			"extraData": {}
		}`)
		var a thehive.OutputAttachment
		require.NoError(t, json.Unmarshal(payload, &a))
		assert.Equal(t, "f.txt", a.Name)
		assert.Nil(t, a.External, "External must be optional for 5.5 compatibility")
	})

	t.Run("OutputProfile decodes without type and forExternal", func(t *testing.T) {
		payload := []byte(`{
			"_id": "~1",
			"_type": "Profile",
			"_createdBy": "admin@thehive.local",
			"_createdAt": 1700000000000,
			"name": "analyst",
			"editable": true,
			"forAdmin": false,
			"forOrg": true,
			"consumesLicense": true
		}`)
		var p thehive.OutputProfile
		require.NoError(t, json.Unmarshal(payload, &p))
		assert.Equal(t, "analyst", p.Name)
		assert.Nil(t, p.Type, "Type must be optional for 5.5 compatibility")
		assert.Nil(t, p.ForExternal, "ForExternal must be optional for 5.5 compatibility")
	})

	t.Run("OutputPublicStatus decodes without imports", func(t *testing.T) {
		payload := []byte(`{
			"sso": false,
			"version": "5.5.14"
		}`)
		var s thehive.OutputPublicStatus
		require.NoError(t, json.Unmarshal(payload, &s))
		assert.Equal(t, "5.5.14", s.Version)
		assert.Nil(t, s.Imports, "Imports must be optional for 5.5 compatibility")
	})
}
