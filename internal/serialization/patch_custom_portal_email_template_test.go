package serialization_test

import (
	"encoding/json"
	"testing"

	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/stretchr/testify/require"
)

var (
	_ json.Marshaler   = components.PatchCustomPortalEmailTemplatePayload{}
	_ json.Unmarshaler = (*components.PatchCustomPortalEmailTemplatePayload)(nil)
)

func TestPatchCustomPortalEmailTemplatePayloadOmitsUnsetEnabled(t *testing.T) {
	subject := "Updated subject"
	update := components.PatchCustomPortalEmailTemplatePayload{
		Content: &components.EmailTemplateContent{Subject: &subject},
	}

	require.Equal(t, map[string]any{
		"content": map[string]any{"subject": subject},
	}, marshalObject(t, update))
}

func TestPatchCustomPortalEmailTemplatePayloadIncludesExplicitEnabled(t *testing.T) {
	for _, enabled := range []bool{false, true} {
		t.Run(boolName(enabled), func(t *testing.T) {
			update := components.PatchCustomPortalEmailTemplatePayload{Enabled: &enabled}

			require.Equal(t, map[string]any{"enabled": enabled}, marshalObject(t, update))
		})
	}
}
