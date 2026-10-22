package regression

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/StrangeBeeCorp/thehive4go/thehive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetActionByEntityIsStillThere pins CortexAPI.GetActionByEntity,
// GET /api/v1/connector/cortex/action/{entityType}/{entityId}. The 6.0
// documentation dropped it, TheHive 6.0 still serves it, and TheHiveMCP reads the
// actions a responder ran through it: scripts/preprocess_openapi.sh puts it back
// from its 5.6.2 definition. Without that step the next regeneration loses it,
// and this test stops compiling.
func TestGetActionByEntityIsStillThere(t *testing.T) {
	var gotMethod, gotPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath = r.Method, r.URL.EscapedPath()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"_createdAt":1791986098031,"_createdBy":"analyst@example.org","_id":"~2416713896","_type":"Action",
			"cortexId":"local","cortexJobId":"0lmxOqEBQp5xTN6g36nT","endDate":1791986103671,
			"objectId":"~2703511728","objectType":"Task","operations":"[]","report":"{}",
			"responderDefinition":"testResponder_1_1_0","responderId":"bdfa784dc600c39e27ca89187a4a1662",
			"responderName":"testResponder_1_1_0","startDate":1791986098130,"status":"Success"}]`))
	}))
	defer srv.Close()

	cfg := thehive.NewConfiguration()
	cfg.Servers = thehive.ServerConfigurations{{URL: srv.URL}}
	client := thehive.NewAPIClient(cfg)

	actions, _, err := client.CortexAPI.GetActionByEntity(context.Background(), "case", "~123").Execute()

	require.NoError(t, err)
	assert.Equal(t, http.MethodGet, gotMethod)
	assert.Equal(t, "/api/v1/connector/cortex/action/case/~123", gotPath)
	require.Len(t, actions, 1)
	assert.Equal(t, "~2416713896", actions[0].UnderscoreId)
}
