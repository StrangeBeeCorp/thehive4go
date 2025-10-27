package integration

import (
	"fmt"
	"testing"
	"time"

	"github.com/StrangeBeeCorp/thehive4go/integration/tests/testutils"
	"github.com/StrangeBeeCorp/thehive4go/thehive"
	"github.com/stretchr/testify/require"
)

func TestAlertOperations(t *testing.T) {
	cfg := testutils.GetTestConfig(t)

	client := testutils.CreateOrgClient(t, cfg)
	ctx := testutils.CreateAuthContext(t, cfg)

	alertTitle := fmt.Sprintf("Integration Test Alert - %d", time.Now().Unix())
	sourceRef := fmt.Sprintf("test-%d", time.Now().Unix())

	createInput := thehive.NewInputCreateAlert("external", "integration-test", sourceRef, alertTitle, "Test alert")
	createInput.SetSeverity(2)
	createInput.SetTlp(2)
	createInput.SetPap(2)
	createInput.SetTags([]string{"integration-test"})

	createResp, httpResp, err := client.AlertAPI.CreateAlert(ctx).InputCreateAlert(*createInput).Execute()
	require.NoError(t, err)
	require.Equal(t, 201, httpResp.StatusCode)

	alertID := createResp.GetUnderscoreId()

	getResp, httpResp, err := client.AlertAPI.GetAlert(ctx, alertID).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, httpResp.StatusCode)
	require.Equal(t, alertTitle, getResp.GetTitle())

	updateInput := thehive.NewInputUpdateAlert()
	newTitle := "Updated Integration Test Alert"
	updateInput.SetTitle(newTitle)

	httpResp, err = client.AlertAPI.UpdateAlert(ctx, alertID).InputUpdateAlert(*updateInput).Execute()
	require.NoError(t, err)
	require.Equal(t, 204, httpResp.StatusCode)

	getResp, httpResp, err = client.AlertAPI.GetAlert(ctx, alertID).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, httpResp.StatusCode)
	require.Equal(t, newTitle, getResp.GetTitle())

	httpResp, err = client.AlertAPI.DeleteAlert(ctx, alertID).Execute()
	require.NoError(t, err)
	require.Equal(t, 204, httpResp.StatusCode)

	_, httpResp, err = client.AlertAPI.GetAlert(ctx, alertID).Execute()
	require.Error(t, err)
	require.Equal(t, 404, httpResp.StatusCode)
}

func TestAlertListing(t *testing.T) {
	cfg := testutils.GetTestConfig(t)

	client := testutils.CreateOrgClient(t, cfg)
	ctx := testutils.CreateAuthContext(t, cfg)

	genericOp := thehive.NewInputQueryGenericOperation("listAlert")
	namedOp := thehive.InputQueryGenericOperationAsInputQueryNamedOperation(genericOp)

	query := thehive.NewInputQuery()
	query.SetQuery([]thehive.InputQueryNamedOperation{namedOp})

	resp, httpResp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(*query).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, httpResp.StatusCode)
	require.NotNil(t, resp)
}
