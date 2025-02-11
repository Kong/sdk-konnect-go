package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"
)

// APIType is an enum for the type of API to use.
type APIType byte

const (
	// GlobalAPI is the global API.
	GlobalAPI APIType = iota
	// RegionalAPI is the regional API.
	RegionalAPI APIType = iota
)

// SDK returns a new SDK instance. It requires the KONNECT_API_PAT and KONNECT_API_URL
// environment variables to be set.
func SDK(t *testing.T, apiType APIType, opts ...func(*sdkkonnectgo.SDK)) *sdkkonnectgo.SDK {
	pat := KonnectPersonalAccessToken(t)
	url := KonnectURL(t)
	sdk := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(
			sdkkonnectcomp.Security{
				PersonalAccessToken: sdkkonnectgo.String(pat),
			},
		),
		sdkkonnectgo.WithServerURL(url),
	)
	for _, opt := range opts {
		opt(sdk)
	}

	if apiType == GlobalAPI {
		var err error
		url, err = replaceFirstSegmentWithGlobal(url)
		require.NoError(t, err)
	}

	require.NotNil(t, sdk)
	return sdk
}
