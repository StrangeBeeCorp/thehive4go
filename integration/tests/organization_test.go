package integration

import (
	"testing"

	"github.com/StrangeBeeCorp/thehive4go/integration/tests/testutils"
	"github.com/StrangeBeeCorp/thehive4go/thehive"
	"github.com/stretchr/testify/require"
)

func TestOrganizationListing(t *testing.T) {
	cfg := testutils.GetTestConfig(t)

	client := testutils.CreateAdminClient(t, cfg)
	ctx := testutils.CreateAuthContext(t, cfg)

	genericOp := thehive.NewInputQueryGenericOperation("listOrganisation")
	namedOp := thehive.InputQueryGenericOperationAsInputQueryNamedOperation(genericOp)

	query := thehive.NewInputQuery()
	query.SetQuery([]thehive.InputQueryNamedOperation{namedOp})

	resp, httpResp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(*query).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, httpResp.StatusCode)
	require.NotNil(t, resp)
}

func TestOrganizationSetup(t *testing.T) {
	cfg := testutils.GetTestConfig(t)

	client := testutils.CreateAdminClient(t, cfg)
	ctx := testutils.CreateAuthContext(t, cfg)

	orgID := testutils.EnsureTestOrganization(t, client, ctx, cfg.OrgName)
	require.NotEmpty(t, orgID)

	testutils.SetupUserPermissions(t, client, ctx, cfg.OrgName)
}
