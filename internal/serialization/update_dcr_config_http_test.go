package serialization_test

import (
	"encoding/json"
	"testing"

	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/stretchr/testify/require"
)

var (
	_ json.Marshaler   = components.UpdateDcrConfigHTTPInRequest{}
	_ json.Unmarshaler = (*components.UpdateDcrConfigHTTPInRequest)(nil)
)

func TestUpdateDcrConfigHTTPInRequestOmitsUnsetAllowMultipleCredentials(t *testing.T) {
	baseURL := "https://example.com/v2/dcr"
	update := components.UpdateDcrConfigHTTPInRequest{DcrBaseURL: &baseURL}

	require.Equal(t, map[string]any{"dcr_base_url": baseURL}, marshalObject(t, update))
}

func TestUpdateDcrConfigHTTPInRequestIncludesExplicitAllowMultipleCredentials(t *testing.T) {
	for _, enabled := range []bool{false, true} {
		t.Run(boolName(enabled), func(t *testing.T) {
			update := components.UpdateDcrConfigHTTPInRequest{AllowMultipleCredentials: &enabled}

			require.Equal(t, map[string]any{"allow_multiple_credentials": enabled}, marshalObject(t, update))
		})
	}
}
