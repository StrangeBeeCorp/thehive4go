package integration

import (
	"encoding/json"
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

func TestAlertListingWithExtraData(t *testing.T) {
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

	// FIX: Check the error
	_, _, err := client.AlertAPI.CreateAlert(ctx).InputCreateAlert(*createInput).Execute()
	require.NoError(t, err, "Failed to create test alert")

	// Create paging operation with extraData
	pageOp := thehive.NewInputQueryPagingOperation(0, 10, "page")
	pageOp.SetExtraData([]string{"importDate", "procedureCount"})
	listOp := thehive.NewInputQueryGenericOperation("listAlert")

	query := thehive.InputQuery{
		Query: []thehive.InputQueryNamedOperation{
			thehive.InputQueryGenericOperationAsInputQueryNamedOperation(listOp),
			thehive.InputQueryPagingOperationAsInputQueryNamedOperation(pageOp),
		},
	}

	resp, httpResp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(query).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, httpResp.StatusCode)
	require.NotNil(t, resp)

	// Convert the untyped response to typed structs
	var alerts []thehive.OutputAlert

	respBytes, err := json.Marshal(resp)
	require.NoError(t, err)

	err = json.Unmarshal(respBytes, &alerts)
	require.NoError(t, err)

	// Check that extraData fields are present in results
	if len(alerts) > 0 {
		firstAlert := alerts[0]
		extraData := firstAlert.GetExtraData()

		_, importDateExists := extraData["importDate"]
		_, procedureCountExists := extraData["procedureCount"]

		require.True(t, importDateExists || procedureCountExists,
			"At least one extraData field should be present, got: %v", extraData)
	}
}
