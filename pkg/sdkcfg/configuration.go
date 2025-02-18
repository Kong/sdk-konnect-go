package sdkcfg

import (
	"os"
)

const (
	// KonnectPersonalAccessTokenEnv is the environment variable name for the Konnect PAT.
	KonnectPersonalAccessTokenEnv = "KONNECT_API_PAT"
	// KonnectURLEnv is the environment variable name for the Konnect URL.
	KonnectURLEnv = "KONNECT_API_URL"
)

// PersonalAccessToken returns the Konnect PAT from the environment.
func PersonalAccessToken() string {
	return os.Getenv(KonnectPersonalAccessTokenEnv)
}

// URL returns the Konnect url from the environment.
func URL() string {
	return os.Getenv(KonnectURLEnv)
}
