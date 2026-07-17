package serialization_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/stretchr/testify/require"
)

var (
	_ json.Marshaler   = components.UpdatePortal{}
	_ json.Unmarshaler = (*components.UpdatePortal)(nil)
)

func TestUpdatePortalOmitsUnsetDefaultedFields(t *testing.T) {
	displayName := "Updated portal"
	update := components.UpdatePortal{DisplayName: &displayName}

	require.Equal(t, map[string]any{"display_name": displayName}, marshalObject(t, update))
}

func TestUpdatePortalIncludesExplicitBooleanValues(t *testing.T) {
	fields := []struct {
		name string
		key  string
		set  func(*components.UpdatePortal, *bool)
	}{
		{
			name: "authentication enabled",
			key:  "authentication_enabled",
			set:  func(update *components.UpdatePortal, value *bool) { update.AuthenticationEnabled = value },
		},
		{
			name: "RBAC enabled",
			key:  "rbac_enabled",
			set:  func(update *components.UpdatePortal, value *bool) { update.RbacEnabled = value },
		},
		{
			name: "SIPR enabled",
			key:  "sipr_enabled",
			set:  func(update *components.UpdatePortal, value *bool) { update.SiprEnabled = value },
		},
		{
			name: "auto approve developers",
			key:  "auto_approve_developers",
			set:  func(update *components.UpdatePortal, value *bool) { update.AutoApproveDevelopers = value },
		},
		{
			name: "auto approve applications",
			key:  "auto_approve_applications",
			set:  func(update *components.UpdatePortal, value *bool) { update.AutoApproveApplications = value },
		},
		{
			name: "developer PII visibility",
			key:  "notifications_developer_pii_visibility_enabled",
			set: func(update *components.UpdatePortal, value *bool) {
				update.NotificationsDeveloperPiiVisibilityEnabled = value
			},
		},
	}

	for _, field := range fields {
		for _, value := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s=%t", field.name, value), func(t *testing.T) {
				update := components.UpdatePortal{}
				field.set(&update, &value)

				require.Equal(t, map[string]any{field.key: value}, marshalObject(t, update))
			})
		}
	}
}

func TestCreatePortalRetainsDefaults(t *testing.T) {
	payload := marshalObject(t, components.CreatePortal{Name: "portal"})

	require.Equal(t, true, payload["authentication_enabled"])
	require.Equal(t, false, payload["rbac_enabled"])
	require.Equal(t, false, payload["sipr_enabled"])
	require.Equal(t, false, payload["auto_approve_developers"])
	require.Equal(t, false, payload["auto_approve_applications"])
	require.Equal(t, false, payload["notifications_developer_pii_visibility_enabled"])
}

func marshalObject(t *testing.T, value any) map[string]any {
	t.Helper()

	data, err := json.Marshal(value)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(data, &payload))
	return payload
}
