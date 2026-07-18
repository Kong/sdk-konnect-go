package serialization_test

import (
	"encoding/json"
	"testing"

	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/stretchr/testify/require"
)

var (
	_ json.Marshaler   = components.UpdatePortalAuditLogWebhook{}
	_ json.Unmarshaler = (*components.UpdatePortalAuditLogWebhook)(nil)
)

func TestUpdatePortalAuditLogWebhookOmitsUnsetEnabled(t *testing.T) {
	destinationID := "11111111-1111-1111-1111-111111111111"
	update := components.UpdatePortalAuditLogWebhook{AuditLogDestinationID: &destinationID}

	require.Equal(t, map[string]any{"audit_log_destination_id": destinationID}, marshalObject(t, update))
}

func TestUpdatePortalAuditLogWebhookIncludesExplicitEnabled(t *testing.T) {
	for _, enabled := range []bool{false, true} {
		t.Run(boolName(enabled), func(t *testing.T) {
			update := components.UpdatePortalAuditLogWebhook{Enabled: &enabled}

			require.Equal(t, map[string]any{"enabled": enabled}, marshalObject(t, update))
		})
	}
}

func boolName(value bool) string {
	if value {
		return "true"
	}
	return "false"
}
