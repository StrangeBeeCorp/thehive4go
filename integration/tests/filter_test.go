package integration

import (
	"encoding/json"
	"testing"

	"github.com/StrangeBeeCorp/thehive4go/integration/tests/testutils"
	"github.com/StrangeBeeCorp/thehive4go/thehive"
	"github.com/stretchr/testify/require"
)

// Helper function to create filters from JSON strings using the generic map approach
// This demonstrates how to work with dynamic/LLM-generated filters
func createFilterFromJSONString(filterString string) (thehive.InputQueryNamedOperation, error) {
	var filterMap map[string]interface{}
	if err := json.Unmarshal([]byte(filterString), &filterMap); err != nil {
		return thehive.InputQueryNamedOperation{}, err
	}

	// Use the generic map converter that's available in the generated client
	return thehive.MapmapOfStringAnyAsInputQueryNamedOperation(&filterMap), nil
}

// TestFilterOperations tests various filter operations using the generated client
func TestFilterOperations(t *testing.T) {
	cfg := testutils.GetTestConfig(t)
	client := testutils.CreateOrgClient(t, cfg)
	ctx := testutils.CreateAuthContext(t, cfg)

	t.Run("TestEqFilterAlerts", func(t *testing.T) {
		// Create a list operation for alerts
		listOp := thehive.NewInputQueryGenericOperation("listAlert")
		listNamedOp := thehive.InputQueryGenericOperationAsInputQueryNamedOperation(listOp)

		// Use the generic map approach for filter - test filtering by alert type
		filterMap := map[string]interface{}{
			"_name": "filter",
			"_eq": map[string]interface{}{
				"_field": "type",
				"_value": "external",
			},
		}
		filterNamedOp := thehive.MapmapOfStringAnyAsInputQueryNamedOperation(&filterMap)

		query := thehive.NewInputQuery()
		query.SetQuery([]thehive.InputQueryNamedOperation{listNamedOp, filterNamedOp})

		resp, httpResp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(*query).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, httpResp.StatusCode)
		require.NotNil(t, resp)
		t.Logf("✅ _eq filter query works for alerts")
	})

	t.Run("TestEqFilterObservables", func(t *testing.T) {
		// Create a list operation for observables
		listOp := thehive.NewInputQueryGenericOperation("listObservable")
		listNamedOp := thehive.InputQueryGenericOperationAsInputQueryNamedOperation(listOp)

		// Use the generic map approach for filter - test filtering by dataType
		filterMap := map[string]interface{}{
			"_name": "filter",
			"_eq": map[string]interface{}{
				"_field": "dataType",
				"_value": "ip",
			},
		}
		filterNamedOp := thehive.MapmapOfStringAnyAsInputQueryNamedOperation(&filterMap)

		query := thehive.NewInputQuery()
		query.SetQuery([]thehive.InputQueryNamedOperation{listNamedOp, filterNamedOp})

		resp, httpResp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(*query).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, httpResp.StatusCode)
		require.NotNil(t, resp)
		t.Logf("✅ _eq filter query works for observables")
	})

	t.Run("TestGtFilter", func(t *testing.T) {
		// Create a list operation for alerts
		listOp := thehive.NewInputQueryGenericOperation("listAlert")
		listNamedOp := thehive.InputQueryGenericOperationAsInputQueryNamedOperation(listOp)

		// Test _gt filter for creation date
		filterMap := map[string]interface{}{
			"_name": "filter",
			"_gt": map[string]interface{}{
				"_field": "_createdAt",
				"_value": 1640995200000, // Unix timestamp
			},
		}
		filterNamedOp := thehive.MapmapOfStringAnyAsInputQueryNamedOperation(&filterMap)

		query := thehive.NewInputQuery()
		query.SetQuery([]thehive.InputQueryNamedOperation{listNamedOp, filterNamedOp})

		resp, httpResp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(*query).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, httpResp.StatusCode)
		require.NotNil(t, resp)
		t.Logf("✅ _gt filter query works")
	})

	t.Run("TestAndFilter", func(t *testing.T) {
		// Create a list operation for alerts
		listOp := thehive.NewInputQueryGenericOperation("listAlert")
		listNamedOp := thehive.InputQueryGenericOperationAsInputQueryNamedOperation(listOp)

		// Test _and filter combining multiple conditions
		filterMap := map[string]interface{}{
			"_name": "filter",
			"_and": []interface{}{
				map[string]interface{}{
					"_gte": map[string]interface{}{
						"_field": "severity",
						"_value": 2,
					},
				},
				map[string]interface{}{
					"_eq": map[string]interface{}{
						"_field": "status",
						"_value": "New",
					},
				},
			},
		}
		filterNamedOp := thehive.MapmapOfStringAnyAsInputQueryNamedOperation(&filterMap)

		query := thehive.NewInputQuery()
		query.SetQuery([]thehive.InputQueryNamedOperation{listNamedOp, filterNamedOp})

		resp, httpResp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(*query).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, httpResp.StatusCode)
		require.NotNil(t, resp)
		t.Logf("✅ _and filter query works")
	})

	t.Run("TestLLMGeneratedFilters", func(t *testing.T) {
		// Test filters generated from JSON strings (simulating LLM output)
		testCases := []struct {
			name       string
			filterJSON string
		}{
			{
				name: "LLM generated equality filter",
				filterJSON: `{
					"_name": "filter",
					"_eq": {
						"_field": "type",
						"_value": "external"
					}
				}`,
			},
			{
				name: "LLM generated range filter",
				filterJSON: `{
					"_name": "filter",
					"_gte": {
						"_field": "severity",
						"_value": 1
					}
				}`,
			},
		}

		// Create a base list operation
		listOp := thehive.NewInputQueryGenericOperation("listAlert")
		listNamedOp := thehive.InputQueryGenericOperationAsInputQueryNamedOperation(listOp)

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				filterNamedOp, err := createFilterFromJSONString(tc.filterJSON)
				require.NoError(t, err)

				query := thehive.NewInputQuery()
				query.SetQuery([]thehive.InputQueryNamedOperation{listNamedOp, filterNamedOp})

				resp, httpResp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(*query).Execute()
				require.NoError(t, err)
				require.Equal(t, 200, httpResp.StatusCode)
				require.NotNil(t, resp)
				t.Logf("✅ LLM generated filter works: %s", tc.name)
			})
		}
	})

	t.Run("TestExtraDataFiltering", func(t *testing.T) {
		// Test that we can list observables and get extraData without separate extraData operation
		listOp := thehive.NewInputQueryGenericOperation("listObservable")
		listNamedOp := thehive.InputQueryGenericOperationAsInputQueryNamedOperation(listOp)

		// Filter on a valid observable field
		filterOp := map[string]interface{}{
			"_name": "filter",
			"_eq": map[string]interface{}{
				"_field": "ioc",
				"_value": true,
			},
		}
		filterNamedOp := thehive.MapmapOfStringAnyAsInputQueryNamedOperation(&filterOp)

		// Add page operation to request extraData (correct way)
		pageOp := map[string]interface{}{
			"_name":     "page",
			"from":      0,
			"to":        10,
			"extraData": []string{"all"},
		}
		pageNamedOp := thehive.MapmapOfStringAnyAsInputQueryNamedOperation(&pageOp)

		query := thehive.NewInputQuery()
		query.SetQuery([]thehive.InputQueryNamedOperation{listNamedOp, filterNamedOp, pageNamedOp})

		resp, httpResp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(*query).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, httpResp.StatusCode)
		require.NotNil(t, resp)
		t.Logf("✅ extraData can be requested via page operation")
	}) // Test building filters entirely from generic maps (LLM use case)
	t.Run("EntirelyGenericMapFilters", func(t *testing.T) {
		// Test the hybrid approach: proper list operation + generic map filter
		// This is what works and is most practical for LLM use
		listOp := thehive.NewInputQueryGenericOperation("listAlert")
		listNamedOp := thehive.InputQueryGenericOperationAsInputQueryNamedOperation(listOp)

		// LLM-generated filter as generic map
		filterMap := map[string]interface{}{
			"_name": "filter",
			"_eq": map[string]interface{}{
				"_field": "type",
				"_value": "external",
			},
		}
		filterNamedOp := thehive.MapmapOfStringAnyAsInputQueryNamedOperation(&filterMap)

		namedOps := []thehive.InputQueryNamedOperation{listNamedOp, filterNamedOp}

		query := thehive.NewInputQuery()
		query.SetQuery(namedOps)

		resp, httpResp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(*query).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, httpResp.StatusCode)
		require.NotNil(t, resp)
		t.Logf("✅ Entirely generic map-based query works - PERFECT FOR LLM USE!")
	})
}
