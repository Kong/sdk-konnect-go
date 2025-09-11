package integration

import (
	"fmt"
	"regexp"
	"strings"
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
	if apiType == GlobalAPI {
		var err error
		url, err = replaceFirstSegmentWithGlobal(url)
		require.NoError(t, err)
	}

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

	require.NotNil(t, sdk)
	return sdk
}

func replaceFirstSegmentWithGlobal(url string) (string, error) {
	if url == "" {
		return "", fmt.Errorf("empty URL provided")
	}

	// Pattern explanation:
	// ^(?:https?:\/\/)? - Optional http:// or https:// at the start
	// ([^\.]+)         - Capture the first segment (anything before the first dot)
	// (\..+)           - Capture everything else starting with the dot
	pattern := `^(?:https?:\/\/)?([^\.]+)(\..+)`
	regex := regexp.MustCompile(pattern)

	matches := regex.FindStringSubmatch(url)
	if matches == nil {
		return "", fmt.Errorf("invalid URL format")
	}

	// matches[0] is the full match
	// matches[1] is the first segment
	// matches[2] is the rest of the URL
	result := "global" + matches[2]

	// If the original URL had the protocol, preserve it
	if strings.HasPrefix(url, "http://") {
		result = "http://" + result
	} else if strings.HasPrefix(url, "https://") {
		result = "https://" + result
	}

	return result, nil
}
