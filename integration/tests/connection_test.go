package integration

import (
	"context"
	"testing"

	"github.com/StrangeBeeCorp/thehive4go/integration/tests/testutils"
	"github.com/stretchr/testify/require"
)

func TestBasicConnection(t *testing.T) {
	cfg := testutils.GetTestConfig(t)

	client := testutils.CreateAdminClient(t, cfg)
	testutils.WaitForTheHive(t, client)

	resp, httpResp, err := client.StatusAPI.GetPlatformPublicStatus(context.Background()).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, httpResp.StatusCode)
	require.NotNil(t, resp)
}

func TestAuthenticatedStatus(t *testing.T) {
	cfg := testutils.GetTestConfig(t)

	client := testutils.CreateAdminClient(t, cfg)
	ctx := testutils.CreateAuthContext(t, cfg)

	resp, httpResp, err := client.StatusAPI.GetPlatformStatus(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, 200, httpResp.StatusCode)
	require.NotNil(t, resp)
}
