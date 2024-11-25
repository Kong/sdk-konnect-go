package integration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	// KonnectPersonalAccessTokenEnv is the environment variable name for the Konnect PAT.
	KonnectPersonalAccessTokenEnv = "KONNECT_API_PAT"
	// KonnectURLEnv is the environment variable name for the Konnect URL.
	KonnectURLEnv = "KONNECT_API_URL"
	// KonnectTestRunIDEnv is the environment variable name for the Konnect test run ID.
	KonnectTestRunIDEnv = "KONNECT_TEST_RUN_ID"
)

// KonnectPersonalAccessToken returns the Konnect PAT from the environment.
func KonnectPersonalAccessToken(t *testing.T) string {
	token := os.Getenv(KonnectPersonalAccessTokenEnv)
	require.NotEmptyf(t, token, "%s is not set", KonnectPersonalAccessTokenEnv)
	return token
}

// KonnectURL returns the Konnect url from the environment.
func KonnectURL(t *testing.T) string {
	url := os.Getenv(KonnectURLEnv)
	require.NotEmptyf(t, url, "%s is not set", KonnectURLEnv)
	return url
}

// KonnectTestRunID returns the Konnect test run ID from the environment.
func KonnectTestRunID(t *testing.T) string {
	id := os.Getenv(KonnectTestRunIDEnv)
	require.NotEmptyf(t, id, "%s is not set", KonnectTestRunIDEnv)
	return id
}
