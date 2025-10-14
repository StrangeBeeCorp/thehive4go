package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
	// Get configuration from environment
	url := getEnv("THEHIVE_URL", "http://thehive:9000")
	username := getEnv("THEHIVE_USERNAME", "admin")
	password := getEnv("THEHIVE_PASSWORD", "secret")

	log.Printf("Starting integration tests against: %s", url)

	// Create TheHive client
	cfg := thehive.NewConfiguration()
	cfg.Host = strings.TrimPrefix(url, "http://")
	cfg.Host = strings.TrimPrefix(cfg.Host, "https://")
	cfg.Scheme = "http"
	if strings.HasPrefix(url, "https://") {
		cfg.Scheme = "https"
	}

	client := thehive.NewAPIClient(cfg)

	// Create authenticated context
	auth := thehive.BasicAuth{
		UserName: username,
		Password: password,
	}
	ctx := context.WithValue(context.Background(), thehive.ContextBasicAuth, auth)

	// Wait for TheHive to be ready
	if err := waitForTheHive(client); err != nil {
		log.Fatalf("TheHive is not ready: %v", err)
	}

	// Run tests
	tests := []struct {
		name string
		test func(*thehive.APIClient, context.Context) error
	}{
		{"Connection", testConnection},
		{"Status API", testStatusAPI},
		{"Alert API", testAlertAPI},
	}

	passed := 0
	failed := 0

	for _, test := range tests {
		log.Printf("Running %s test...", test.name)
		if err := test.test(client, ctx); err != nil {
			log.Printf("❌ %s FAILED: %v", test.name, err)
			failed++
		} else {
			log.Printf("✅ %s PASSED", test.name)
			passed++
		}
	}

	// Print summary
	log.Printf("\n=== RESULTS ===")
	log.Printf("Passed: %d", passed)
	log.Printf("Failed: %d", failed)
	log.Printf("Total: %d", passed+failed)

	if failed > 0 {
		os.Exit(1)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func waitForTheHive(client *thehive.APIClient) error {
	for i := 0; i < 30; i++ {
		resp, httpResp, err := client.StatusAPI.GetPlatformPublicStatus(context.Background()).Execute()
		if err == nil && httpResp.StatusCode == 200 && resp != nil {
			log.Printf("TheHive is ready! Version: %s", resp.GetVersion())
			return nil
		}
		log.Printf("Waiting for TheHive... (attempt %d/30)", i+1)
		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("TheHive not ready after 30 attempts")
}

// testConnection verifies basic connectivity to TheHive
func testConnection(client *thehive.APIClient, ctx context.Context) error {
	resp, httpResp, err := client.StatusAPI.GetPlatformPublicStatus(context.Background()).Execute()
	if err != nil {
		return fmt.Errorf("failed to connect to TheHive: %v", err)
	}

	if httpResp.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", httpResp.StatusCode)
	}

	if resp.GetVersion() == "" {
		return fmt.Errorf("received empty version")
	}

	log.Printf("Successfully connected to TheHive version: %s", resp.GetVersion())
	return nil
}

// testStatusAPI tests the status API endpoints
func testStatusAPI(client *thehive.APIClient, ctx context.Context) error {
	// Test public status first (no auth needed)
	resp, httpResp, err := client.StatusAPI.GetPlatformPublicStatus(context.Background()).Execute()
	if err != nil {
		return fmt.Errorf("GetPlatformPublicStatus failed: %v", err)
	}
	if httpResp.StatusCode != 200 {
		return fmt.Errorf("expected status 200, got %d", httpResp.StatusCode)
	}
	if resp.GetVersion() == "" {
		return fmt.Errorf("version is empty")
	}

	// Test authenticated status - let it fail if there's an int32/int64 issue
	authResp, httpResp, err := client.StatusAPI.GetPlatformStatus(ctx).Execute()
	if err != nil {
		return fmt.Errorf("GetPlatformStatus failed: %v", err)
	}
	if httpResp.StatusCode != 200 {
		return fmt.Errorf("expected status 200, got %d", httpResp.StatusCode)
	}
	if authResp.GetVersion() == "" {
		return fmt.Errorf("authenticated version is empty")
	}

	return nil
}

// testAlertAPI tests the alert API endpoints
func testAlertAPI(client *thehive.APIClient, ctx context.Context) error {
	// Create a simple alert
	title := fmt.Sprintf("Integration Test Alert - %d", time.Now().Unix())
	description := "Test alert created by integration tests"

	// NewInputCreateAlert requires: type_, source, sourceRef, title, description
	alertInput := thehive.NewInputCreateAlert("integration-test", "integration", "test-ref-"+fmt.Sprintf("%d", time.Now().Unix()), title, description)
	alertInput.SetSeverity(2)
	alertInput.SetTlp(2)
	alertInput.SetPap(2)
	alertInput.SetTags([]string{"integration-test"})

	createdAlert, httpResp, err := client.AlertAPI.CreateAlert(ctx).InputCreateAlert(*alertInput).Execute()
	if err != nil {
		return fmt.Errorf("CreateAlert failed: %v", err)
	}
	if httpResp.StatusCode != 201 {
		return fmt.Errorf("expected status 201, got %d", httpResp.StatusCode)
	}
	if createdAlert.GetUnderscoreId() == "" {
		return fmt.Errorf("created alert ID is empty")
	}

	alertID := createdAlert.GetUnderscoreId()
	log.Printf("Created alert with ID: %s", alertID)

	// Get the alert back
	retrievedAlert, httpResp, err := client.AlertAPI.GetAlert(ctx, alertID).Execute()
	if err != nil {
		return fmt.Errorf("GetAlert failed: %v", err)
	}
	if httpResp.StatusCode != 200 {
		return fmt.Errorf("expected status 200, got %d", httpResp.StatusCode)
	}
	if retrievedAlert.GetUnderscoreId() != alertID {
		return fmt.Errorf("retrieved alert ID mismatch: expected %s, got %s", alertID, retrievedAlert.GetUnderscoreId())
	}

	// Clean up - delete the alert
	_, err = client.AlertAPI.DeleteAlert(ctx, alertID).Execute()
	if err != nil {
		log.Printf("Warning: failed to cleanup alert %s: %v", alertID, err)
	} else {
		log.Printf("Cleaned up alert: %s", alertID)
	}

	return nil
}
