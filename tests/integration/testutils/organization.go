package testutils

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/StrangeBeeCorp/thehive4go/thehive"
	"github.com/stretchr/testify/require"
)

func EnsureTestOrganization(t *testing.T, client *thehive.APIClient, ctx context.Context, orgName string) string {
	t.Helper()

	genericOp := thehive.NewInputQueryGenericOperation("listOrganisation")
	namedOp := thehive.InputQueryGenericOperationAsInputQueryNamedOperation(genericOp)

	query := thehive.NewInputQuery()
	query.SetQuery([]thehive.InputQueryNamedOperation{namedOp})

	resp, httpResp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(*query).Execute()
	var orgs []thehive.OutputOrganisation
	if err == nil && httpResp.StatusCode == 200 && resp != nil {
		jsonBytes, _ := json.Marshal(resp)
		if err := json.Unmarshal(jsonBytes, &orgs); err == nil {
			for _, org := range orgs {
				if org.GetName() == orgName {
					return org.GetUnderscoreId()
				}
			}
		}
	}

	createOrgInput := thehive.NewInputCreateOrganisation(orgName, "Integration test organization")
	createResp, httpResp, err := client.OrganisationAPI.CreateOrganisation(ctx).InputCreateOrganisation(*createOrgInput).Execute()
	if err != nil && httpResp != nil && (httpResp.StatusCode == 409 || httpResp.StatusCode == 403) {
		return orgName
	}
	require.NoError(t, err)
	require.Equal(t, 201, httpResp.StatusCode)

	return createResp.GetUnderscoreId()
}

func ListOrganizations(t *testing.T, client *thehive.APIClient, ctx context.Context) []thehive.OutputOrganisation {
	t.Helper()

	genericOp := thehive.NewInputQueryGenericOperation("listOrganisation")
	namedOp := thehive.InputQueryGenericOperationAsInputQueryNamedOperation(genericOp)

	query := thehive.NewInputQuery()
	query.SetQuery([]thehive.InputQueryNamedOperation{namedOp})

	resp, httpResp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(*query).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, httpResp.StatusCode)

	var orgs []thehive.OutputOrganisation
	jsonBytes, _ := json.Marshal(resp)
	if err := json.Unmarshal(jsonBytes, &orgs); err != nil {
		t.Fatalf("Failed to unmarshal organizations: %v", err)
	}

	return orgs
}

func GetOrganizationByName(t *testing.T, organizations []thehive.OutputOrganisation, name string) *thehive.OutputOrganisation {
	t.Helper()

	for i, org := range organizations {
		if org.GetName() == name {
			return &organizations[i]
		}
	}

	t.Fatalf("Organization '%s' not found", name)
	return nil
}

func SetupUserPermissions(t *testing.T, client *thehive.APIClient, ctx context.Context, orgName string) {
	t.Helper()

	userResp, httpResp, err := client.UserAPI.GetCurrentUserInfo(ctx).Execute()
	if err != nil || httpResp.StatusCode != http.StatusOK || userResp == nil {
		t.Fatalf("Could not get current user info: %v", err)
		return
	}

	userID := userResp.GetUnderscoreId()
	t.Logf("Current user ID: %s", userID)

	// Try to set user organization assignments
	orgAssignments := []thehive.InputUserOrganisation{
		{
			Organisation: orgName,
			Profile:      "org-admin",
		},
		{
			Organisation: "admin",
			Profile:      "admin",
		},
	}

	updateInput := thehive.NewInputSetUserOrganisations()
	updateInput.SetOrganisations(orgAssignments)

	_, httpResp, err = client.UserAPI.SetUserOrganisations(ctx, userID).InputSetUserOrganisations(*updateInput).Execute()

	if err != nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("Failed to set user organizations: %v, got status: %d", err, httpResp.StatusCode)
		return
	}
	t.Logf("Successfully assigned user to organizations")
}
