package serialization_test

import (
	"encoding/json"
	"testing"

	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/stretchr/testify/require"
)

var (
	_ json.Marshaler   = components.UpdateIdentityProvider{}
	_ json.Unmarshaler = (*components.UpdateIdentityProvider)(nil)
	_ json.Marshaler   = components.PortalUpdateIdentityProvider{}
	_ json.Unmarshaler = (*components.PortalUpdateIdentityProvider)(nil)
)

func TestUpdateIdentityProviderOmitsUnsetEnabled(t *testing.T) {
	loginPath := "oidc-login"
	update := components.UpdateIdentityProvider{LoginPath: &loginPath}

	require.Equal(t, map[string]any{"login_path": loginPath}, marshalObject(t, update))
}

func TestPortalUpdateIdentityProviderOmitsUnsetEnabled(t *testing.T) {
	config := components.CreatePortalUpdateIdentityProviderConfigOIDCIdentityProviderConfig(
		components.OIDCIdentityProviderConfig{
			IssuerURL: "https://idp.example.test",
			ClientID:  "client-id",
		},
	)
	update := components.PortalUpdateIdentityProvider{Config: &config}

	require.Equal(t, map[string]any{
		"config": map[string]any{
			"issuer_url": "https://idp.example.test",
			"client_id":  "client-id",
		},
	}, marshalObject(t, update))
}

func TestPortalIdentityProviderUpdatesIncludeExplicitEnabled(t *testing.T) {
	for _, enabled := range []bool{false, true} {
		t.Run(boolName(enabled), func(t *testing.T) {
			require.Equal(t, map[string]any{"enabled": enabled}, marshalObject(t,
				components.UpdateIdentityProvider{Enabled: &enabled}))
			require.Equal(t, map[string]any{"enabled": enabled}, marshalObject(t,
				components.PortalUpdateIdentityProvider{Enabled: &enabled}))
		})
	}
}
