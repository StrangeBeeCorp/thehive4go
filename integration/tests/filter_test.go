package integration

import (
	"encoding/json"
	"testing"

	"github.com/StrangeBeeCorp/thehive4go/integration/tests/testutils"
	"github.com/StrangeBeeCorp/thehive4go/thehive"
	"github.com/stretchr/testify/require"
)

func TestInputQueryFilterOperationEnhancements(t *testing.T) {
	// This test verifies that the InputQueryFilterOperation struct now supports
	// all filter operations, not just _eq

	t.Run("EqualityFilter", func(t *testing.T) {
		filterOp := thehive.NewInputQueryFilterOperation("filter")
		filterOp.SetEq(map[string]interface{}{
			"_field": "type",
			"_value": "external",
		})

		// Verify it serializes correctly
		data, err := json.Marshal(filterOp)
		require.NoError(t, err)
		require.Contains(t, string(data), `"_eq"`)
		require.Contains(t, string(data), `"_name":"filter"`)
		t.Logf("✅ Equality filter works: %s", string(data))
	})

	t.Run("LikeFilter", func(t *testing.T) {
		filterOp := thehive.NewInputQueryFilterOperation("filter")
		filterOp.SetLike(thehive.FieldValue{
			Field: "title",
			Value: "test*",
		})

		// Verify it serializes correctly
		data, err := json.Marshal(filterOp)
		require.NoError(t, err)
		require.Contains(t, string(data), `"_like"`)
		require.Contains(t, string(data), `"_name":"filter"`)
		t.Logf("✅ Like filter works: %s", string(data))
	})

	t.Run("GreaterThanFilter", func(t *testing.T) {
		filterOp := thehive.NewInputQueryFilterOperation("filter")
		filterOp.SetGt(thehive.FieldValue{
			Field: "severity",
			Value: 1,
		})

		// Verify it serializes correctly
		data, err := json.Marshal(filterOp)
		require.NoError(t, err)
		require.Contains(t, string(data), `"_gt"`)
		require.Contains(t, string(data), `"_name":"filter"`)
		t.Logf("✅ Greater than filter works: %s", string(data))
	})

	t.Run("AndOrFilters", func(t *testing.T) {
		// Test _and filter combining multiple conditions
		filterOp := thehive.NewInputQueryFilterOperation("filter")

		// Create individual filters
		titleFilter := thehive.EqAsFilter(thehive.NewEq(thehive.FieldValue{
			Field: "type",
			Value: "external",
		}))

		severityFilter := thehive.GteAsFilter(thehive.NewGte(thehive.FieldValue{
			Field: "severity",
			Value: 2,
		}))

		andFilters := []thehive.Filter{titleFilter, severityFilter}
		filterOp.SetAnd(andFilters)

		// Verify it serializes correctly
		data, err := json.Marshal(filterOp)
		require.NoError(t, err)
		require.Contains(t, string(data), `"_and"`)
		require.Contains(t, string(data), `"_name":"filter"`)
		t.Logf("✅ AND filter works: %s", string(data))

		// Test _or filter
		filterOp2 := thehive.NewInputQueryFilterOperation("filter")
		orFilters := []thehive.Filter{titleFilter, severityFilter}
		filterOp2.SetOr(orFilters)

		data2, err := json.Marshal(filterOp2)
		require.NoError(t, err)
		require.Contains(t, string(data2), `"_or"`)
		require.Contains(t, string(data2), `"_name":"filter"`)
		t.Logf("✅ OR filter works: %s", string(data2))
	})

	t.Run("AllFilterMethodsAvailable", func(t *testing.T) {
		// Test that all filter operation methods are available
		filterOp := thehive.NewInputQueryFilterOperation("filter")

		// These would fail to compile if the methods don't exist
		require.NotNil(t, filterOp.SetEq)
		require.NotNil(t, filterOp.SetAnd)
		require.NotNil(t, filterOp.SetOr)
		require.NotNil(t, filterOp.SetNot)
		require.NotNil(t, filterOp.SetLike)
		require.NotNil(t, filterOp.SetGt)
		require.NotNil(t, filterOp.SetGte)
		require.NotNil(t, filterOp.SetLt)
		require.NotNil(t, filterOp.SetLte)
		require.NotNil(t, filterOp.SetNe)
		require.NotNil(t, filterOp.SetIn)
		require.NotNil(t, filterOp.SetBetween)
		require.NotNil(t, filterOp.SetContains)
		require.NotNil(t, filterOp.SetEndsWith)
		require.NotNil(t, filterOp.SetStartsWith)
		require.NotNil(t, filterOp.SetMatch)
		require.NotNil(t, filterOp.SetAny)
		require.NotNil(t, filterOp.SetUnderscoreId)

		t.Logf("✅ All filter operations are available (total: 18 operations)")
	})
}

func TestFilterOperationIntegration(t *testing.T) {
	cfg := testutils.GetTestConfig(t)
	client := testutils.CreateOrgClient(t, cfg)
	ctx := testutils.CreateAuthContext(t, cfg)

	t.Run("EqualityFilterQuery", func(t *testing.T) {
		// Test that queries with the _eq filter work (this was already working before)
		listOp := thehive.NewInputQueryGenericOperation("listAlert")
		listNamedOp := thehive.InputQueryGenericOperationAsInputQueryNamedOperation(listOp)

		filterOp := thehive.NewInputQueryFilterOperation("filter")
		filterOp.SetEq(map[string]interface{}{
			"_field": "type",
			"_value": "external",
		})
		filterNamedOp := thehive.InputQueryFilterOperationAsInputQueryNamedOperation(filterOp)

		query := thehive.NewInputQuery()
		query.SetQuery([]thehive.InputQueryNamedOperation{listNamedOp, filterNamedOp})

		resp, httpResp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(*query).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, httpResp.StatusCode)
		require.NotNil(t, resp)
		t.Logf("✅ _eq filter query works")
	})

	t.Run("GreaterThanFilterQuery", func(t *testing.T) {
		// Test that queries with the _gt filter work (this was NOT working before our fix)
		listOp := thehive.NewInputQueryGenericOperation("listAlert")
		listNamedOp := thehive.InputQueryGenericOperationAsInputQueryNamedOperation(listOp)

		filterOp := thehive.NewInputQueryFilterOperation("filter")
		filterOp.SetGt(thehive.FieldValue{
			Field: "severity",
			Value: 0, // Find alerts with severity > 0
		})
		filterNamedOp := thehive.InputQueryFilterOperationAsInputQueryNamedOperation(filterOp)

		query := thehive.NewInputQuery()
		query.SetQuery([]thehive.InputQueryNamedOperation{listNamedOp, filterNamedOp})

		resp, httpResp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(*query).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, httpResp.StatusCode)
		require.NotNil(t, resp)
		t.Logf("✅ _gt filter query works (this was broken before our fix)")
	})

	t.Run("AndFilterQuery", func(t *testing.T) {
		// Test that queries with the _and filter work (this was NOT working before our fix)
		listOp := thehive.NewInputQueryGenericOperation("listAlert")
		listNamedOp := thehive.InputQueryGenericOperationAsInputQueryNamedOperation(listOp)

		// Create individual filters for the AND operation
		typeFilter := thehive.EqAsFilter(thehive.NewEq(thehive.FieldValue{
			Field: "type",
			Value: "external",
		}))

		severityFilter := thehive.GteAsFilter(thehive.NewGte(thehive.FieldValue{
			Field: "severity",
			Value: 1,
		}))

		filterOp := thehive.NewInputQueryFilterOperation("filter")
		filterOp.SetAnd([]thehive.Filter{typeFilter, severityFilter})
		filterNamedOp := thehive.InputQueryFilterOperationAsInputQueryNamedOperation(filterOp)

		query := thehive.NewInputQuery()
		query.SetQuery([]thehive.InputQueryNamedOperation{listNamedOp, filterNamedOp})

		resp, httpResp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(*query).Execute()
		require.NoError(t, err)
		require.Equal(t, 200, httpResp.StatusCode)
		require.NotNil(t, resp)
		t.Logf("✅ _and filter query works (this was broken before our fix)")
	})
}
