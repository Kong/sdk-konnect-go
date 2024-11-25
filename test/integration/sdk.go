package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"
)

// SDK returns a new SDK instance. It requires the KONNECT_API_PAT and KONNECT_API_URL
// environment variables to be set.
func SDK(t *testing.T) *sdkkonnectgo.SDK {
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
	require.NotNil(t, sdk)
	return sdk
}
