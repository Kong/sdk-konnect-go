// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ConfigureOIDCIdentityProviderConfig - The identity provider that contains configuration data for the OIDC authentication integration.
type ConfigureOIDCIdentityProviderConfig struct {
	// The issuer URI of the identity provider. This is the URL where the provider's metadata can be obtained.
	IssuerURL string `json:"issuer_url"`
	// The client ID assigned to your application by the identity provider.
	ClientID string `json:"client_id"`
	// The Client Secret assigned to your application by the identity provider.
	ClientSecret *string `json:"client_secret,omitempty"`
	// The scopes requested by your application when authenticating with the identity provider.
	Scopes []string `json:"scopes,omitempty"`
	// Defines the mappings between OpenID Connect (OIDC) claims and local claims used by your application for
	// authentication.
	//
	ClaimMappings *SchemasOIDCIdentityProviderClaimMappings `json:"claim_mappings,omitempty"`
}

func (o *ConfigureOIDCIdentityProviderConfig) GetIssuerURL() string {
	if o == nil {
		return ""
	}
	return o.IssuerURL
}

func (o *ConfigureOIDCIdentityProviderConfig) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *ConfigureOIDCIdentityProviderConfig) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

func (o *ConfigureOIDCIdentityProviderConfig) GetScopes() []string {
	if o == nil {
		return nil
	}
	return o.Scopes
}

func (o *ConfigureOIDCIdentityProviderConfig) GetClaimMappings() *SchemasOIDCIdentityProviderClaimMappings {
	if o == nil {
		return nil
	}
	return o.ClaimMappings
}
