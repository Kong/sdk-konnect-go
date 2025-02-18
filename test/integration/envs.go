package integration

import (
	"os"
	"testing"

	"github.com/Kong/sdk-konnect-go/pkg/sdkcfg"
	"github.com/stretchr/testify/require"
)

const (
	// KonnectTestRunIDEnv is the environment variable name for the Konnect test run ID.
	KonnectTestRunIDEnv = "KONNECT_TEST_RUN_ID"
)

// KonnectPersonalAccessToken returns the Konnect PAT from the environment.
func KonnectPersonalAccessToken(t *testing.T) string {
	token := sdkcfg.PersonalAccessToken()
	require.NotEmpty(t, token, "Personal Access Token is not set")
	return token
}

// KonnectURL returns the Konnect url from the environment.
func KonnectURL(t *testing.T) string {
	url := sdkcfg.URL()
	require.NotEmpty(t, url, "Konnect API URL is not set")
	return url
}

// KonnectTestRunID returns the Konnect test run ID from the environment.
func KonnectTestRunID(t *testing.T) string {
	id := os.Getenv(KonnectTestRunIDEnv)
	require.NotEmptyf(t, id, "%s is not set", KonnectTestRunIDEnv)
	return id
}
