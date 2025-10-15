package integration

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/StrangeBee/TheHive4Go/thehive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	defaultTimeout = 30 * time.Second
)

// TestConfig holds configuration for integration tests
type TestConfig struct {
	URL      string
	Username string
	Password string
	OrgName  string
}

// getTestConfig retrieves test configuration from environment variables
func getTestConfig(t *testing.T) *TestConfig {
	t.Helper()

	cfg := &TestConfig{
		URL:      getEnv("THEHIVE_URL", "http://thehive:9000"),
		Username: getEnv("THEHIVE_USERNAME", "admin@thehive.local"),
		Password: getEnv("THEHIVE_PASSWORD", "secret"),
		OrgName:  getEnv("THEHIVE_ORG", "main-org"),
	}

	t.Logf("Test configuration: URL=%s, Username=%s, Org=%s", cfg.URL, cfg.Username, cfg.OrgName)
	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// createAdminClient creates a client with admin credentials
func createAdminClient(t *testing.T, cfg *TestConfig) *thehive.ClientWithResponses {
	t.Helper()

	client, err := thehive.NewClientWithResponses(cfg.URL, thehive.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.SetBasicAuth(cfg.Username, cfg.Password)
		return nil
	}))

	require.NoError(t, err, "Failed to create admin client")
	return client
}

// createOrgClient creates a client with organization context
func createOrgClient(t *testing.T, cfg *TestConfig) *thehive.ClientWithResponses {
	t.Helper()

	client, err := thehive.NewClientWithResponses(cfg.URL, thehive.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.SetBasicAuth(cfg.Username, cfg.Password)
		req.Header.Set("X-Organisation", cfg.OrgName)
		return nil
	}))

	require.NoError(t, err, "Failed to create org client")
	return client
}

// listOrganizations lists all organizations using manual JSON query (workaround for union type bug)
func listOrganizations(t *testing.T, client *thehive.ClientWithResponses) []thehive.OutputOrganisation {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// Manual JSON query to workaround union type serialization bug
	manualJSON := `{"query": [{"_name": "listOrganisation"}]}`
	resp, err := client.QueryAPIWithBody(ctx, "application/json", strings.NewReader(manualJSON))
	require.NoError(t, err, "Failed to query organizations")
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "Organizations query failed")

	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Failed to read response body")

	var orgs []thehive.OutputOrganisation
	err = json.Unmarshal(bodyBytes, &orgs)
	require.NoError(t, err, "Failed to parse organizations")

	t.Logf("Found %d organizations", len(orgs))
	for _, org := range orgs {
		t.Logf("  - %s (ID: %s)", org.Name, org.UnderscoreId)
	}

	return orgs
}

// listAlerts lists all alerts using manual JSON query (workaround for union type bug)
func listAlerts(t *testing.T, client *thehive.ClientWithResponses) []thehive.OutputAlert {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// Manual JSON query to workaround union type serialization bug
	manualJSON := `{"query": [{"_name": "listAlert"}]}`
	resp, err := client.QueryAPIWithBody(ctx, "application/json", strings.NewReader(manualJSON))
	require.NoError(t, err, "Failed to query alerts")
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "Alerts query failed")

	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Failed to read response body")

	var alerts []thehive.OutputAlert
	err = json.Unmarshal(bodyBytes, &alerts)
	require.NoError(t, err, "Failed to parse alerts")

	t.Logf("Found %d alerts", len(alerts))

	return alerts
}

// ensureOrganizationExists ensures the test organization exists, creates if not
func ensureOrganizationExists(t *testing.T, adminClient *thehive.ClientWithResponses, orgName string) string {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	orgs := listOrganizations(t, adminClient)

	// Check if organization exists
	for _, org := range orgs {
		if org.Name == orgName {
			t.Logf("Organization '%s' already exists with ID: %s", orgName, org.UnderscoreId)
			return org.UnderscoreId
		}
	}

	// Create organization if it doesn't exist
	t.Logf("Creating organization '%s'", orgName)
	resp, err := adminClient.CreateOrganisationWithResponse(ctx, thehive.InputCreateOrganisation{
		Name:        orgName,
		Description: "Integration test organization",
	})

	require.NoError(t, err, "Failed to create organization")
	require.Equal(t, http.StatusCreated, resp.StatusCode(), "Organization creation failed")
	require.NotNil(t, resp.JSON201, "Organization creation response is nil")

	orgID := resp.JSON201.UnderscoreId
	t.Logf("Created organization '%s' with ID: %s", orgName, orgID)
	return orgID
}

// setupUserOrganizations assigns the admin user to the test organization
func setupUserOrganizations(t *testing.T, adminClient *thehive.ClientWithResponses, orgName string) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// Get current user info
	userResp, err := adminClient.GetCurrentUserInfoWithResponse(ctx)
	require.NoError(t, err, "Failed to get current user")
	require.Equal(t, http.StatusOK, userResp.StatusCode(), "Get user info failed")
	require.NotNil(t, userResp.JSON200, "User info response is nil")

	userID := userResp.JSON200.UnderscoreId
	t.Logf("Current user ID: %s", userID)

	// Set user organization assignments
	orgAssignments := []thehive.InputUserOrganisation{
		{
			Organisation: orgName,
			Profile:      "org-admin",
			Default:      &[]bool{true}[0],
		},
		{
			Organisation: "admin",
			Profile:      "admin",
		},
	}

	updateResp, err := adminClient.SetUserOrganisationsWithResponse(ctx, userID, thehive.InputSetUserOrganisations{
		Organisations: &orgAssignments,
	})

	require.NoError(t, err, "Failed to set user organizations")
	require.Equal(t, http.StatusOK, updateResp.StatusCode(), "Set user organizations failed")

	t.Logf("Successfully assigned user to organizations")
}

func TestTheHiveIntegration(t *testing.T) {
	cfg := getTestConfig(t)
	adminClient := createAdminClient(t, cfg)

	t.Run("ListOrganizations", func(t *testing.T) {
		orgs := listOrganizations(t, adminClient)
		assert.NotEmpty(t, orgs, "Should have at least one organization")
	})

	t.Run("EnsureTestOrganization", func(t *testing.T) {
		orgID := ensureOrganizationExists(t, adminClient, cfg.OrgName)
		assert.NotEmpty(t, orgID, "Organization ID should not be empty")
	})

	t.Run("SetupUserPermissions", func(t *testing.T) {
		setupUserOrganizations(t, adminClient, cfg.OrgName)
		// Test passes if no error occurs
	})

	t.Run("AlertOperations", func(t *testing.T) {
		// Ensure prerequisites
		_ = ensureOrganizationExists(t, adminClient, cfg.OrgName)
		setupUserOrganizations(t, adminClient, cfg.OrgName)

		// Create organization-specific client
		orgClient := createOrgClient(t, cfg)

		t.Run("CreateAlert", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()

			severity := thehive.InputCreateAlertSeverityN3
			tlp := thehive.InputCreateAlertTlpN2
			alert := thehive.InputCreateAlert{
				Title:       "Integration Test Alert",
				Description: "Alert created during integration testing",
				Type:        "external",
				Source:      "integration-test",
				SourceRef:   "test-001",
				Severity:    &severity,
				Tlp:         &tlp,
			}

			resp, err := orgClient.CreateAlertWithResponse(ctx, alert)
			require.NoError(t, err, "Failed to create alert")
			require.Equal(t, http.StatusCreated, resp.StatusCode(), "Alert creation failed")
			require.NotNil(t, resp.JSON201, "Alert creation response is nil")

			alertID := resp.JSON201.UnderscoreId
			assert.NotEmpty(t, alertID, "Alert ID should not be empty")
			assert.Equal(t, alert.Title, resp.JSON201.Title, "Alert title mismatch")

			t.Logf("Created alert with ID: %s", alertID)

			// Cleanup: Delete the created alert
			t.Cleanup(func() {
				cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), defaultTimeout)
				defer cleanupCancel()

				deleteResp, err := orgClient.DeleteAlertWithResponse(cleanupCtx, alertID)
				if err != nil {
					t.Logf("Failed to cleanup alert %s: %v", alertID, err)
				} else if deleteResp.StatusCode() != http.StatusNoContent {
					t.Logf("Alert cleanup returned status: %d", deleteResp.StatusCode())
				} else {
					t.Logf("Successfully cleaned up alert %s", alertID)
				}
			})
		})

		t.Run("ListAlerts", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
			defer cancel()

			// Create a test alert first
			severity := thehive.InputCreateAlertSeverityN2
			tlp := thehive.InputCreateAlertTlpN1
			alert := thehive.InputCreateAlert{
				Title:       "List Test Alert",
				Description: "Alert for testing list functionality",
				Type:        "external",
				Source:      "integration-test",
				SourceRef:   "list-test-001",
				Severity:    &severity,
				Tlp:         &tlp,
			}

			createResp, err := orgClient.CreateAlertWithResponse(ctx, alert)
			require.NoError(t, err, "Failed to create test alert")
			require.Equal(t, http.StatusCreated, createResp.StatusCode())
			require.NotNil(t, createResp.JSON201)

			alertID := createResp.JSON201.UnderscoreId

			// Cleanup
			t.Cleanup(func() {
				cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), defaultTimeout)
				defer cleanupCancel()
				orgClient.DeleteAlertWithResponse(cleanupCtx, alertID)
			})

			// Test listing alerts
			alerts := listAlerts(t, orgClient)
			assert.NotEmpty(t, alerts, "Should have at least one alert")

			// Find our created alert in the list
			found := false
			for _, a := range alerts {
				if a.UnderscoreId == alertID {
					found = true
					assert.Equal(t, alert.Title, a.Title, "Alert title mismatch in list")
					break
				}
			}
			assert.True(t, found, "Created alert should be in the list")

			t.Logf("Successfully listed %d alerts", len(alerts))
		})
	})
}

// TestTheHiveConnection tests basic connectivity to TheHive
func TestTheHiveConnection(t *testing.T) {
	cfg := getTestConfig(t)

	_, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// Test basic HTTP connectivity
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(cfg.URL + "/api/status")
	require.NoError(t, err, "Failed to connect to TheHive")
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "TheHive status endpoint should return 200")

	t.Logf("Successfully connected to TheHive at %s", cfg.URL)
}
