package testutils

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/StrangeBeeCorp/thehive4go/thehive"
)

// Config holds configuration for integration tests
type Config struct {
	URL      string
	Username string
	Password string
	OrgName  string
}

func GetTestConfig(t *testing.T) *Config {
	t.Helper()

	return &Config{
		URL:      getEnv("THEHIVE_URL", "http://thehive:9000"),
		Username: getEnv("THEHIVE_USERNAME", "admin@thehive.local"),
		Password: getEnv("THEHIVE_PASSWORD", "secret"),
		OrgName:  getEnv("THEHIVE_ORG", "main-org"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func CreateAdminClient(t *testing.T, cfg *Config) *thehive.APIClient {
	t.Helper()

	clientCfg := thehive.NewConfiguration()
	clientCfg.Host = strings.TrimPrefix(cfg.URL, "http://")
	clientCfg.Host = strings.TrimPrefix(clientCfg.Host, "https://")
	clientCfg.Scheme = "http"
	if strings.HasPrefix(cfg.URL, "https://") {
		clientCfg.Scheme = "https"
	}

	return thehive.NewAPIClient(clientCfg)
}

func CreateOrgClient(t *testing.T, cfg *Config) *thehive.APIClient {
	t.Helper()

	// First, create an admin client to set up the organization and permissions
	adminClient := CreateAdminClient(t, cfg)
	ctx := CreateAuthContext(t, cfg)

	// Ensure the organization exists and permissions are set using admin client
	EnsureTestOrganization(t, adminClient, ctx, cfg.OrgName)
	SetupUserPermissions(t, adminClient, ctx, cfg.OrgName)

	// Now create the organization-specific client
	clientCfg := thehive.NewConfiguration()
	clientCfg.Host = strings.TrimPrefix(cfg.URL, "http://")
	clientCfg.Host = strings.TrimPrefix(clientCfg.Host, "https://")
	clientCfg.Scheme = "http"
	if strings.HasPrefix(cfg.URL, "https://") {
		clientCfg.Scheme = "https"
	}

	clientCfg.AddDefaultHeader("X-Organisation", cfg.OrgName)

	return thehive.NewAPIClient(clientCfg)
}

func CreateAuthContext(t *testing.T, cfg *Config) context.Context {
	t.Helper()

	auth := thehive.BasicAuth{
		UserName: cfg.Username,
		Password: cfg.Password,
	}
	return context.WithValue(context.Background(), thehive.ContextBasicAuth, auth)
}

func WaitForTheHive(t *testing.T, client *thehive.APIClient) {
	t.Helper()

	for i := 0; i < 30; i++ {
		resp, httpResp, err := client.StatusAPI.GetPlatformPublicStatus(context.Background()).Execute()
		if err == nil && httpResp.StatusCode == 200 && resp != nil {
			return
		}
		time.Sleep(2 * time.Second)
	}
	t.Fatal("TheHive not ready after 30 attempts")
}
