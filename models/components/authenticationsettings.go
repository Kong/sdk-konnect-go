// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// AuthenticationSettings - Response for authentication settings endpoint
type AuthenticationSettings struct {
	// The organization has basic auth enabled.
	BasicAuthEnabled *bool `json:"basic_auth_enabled,omitempty"`
	// The organization has OIDC disabled.
	OidcAuthEnabled *bool `json:"oidc_auth_enabled,omitempty"`
	// The organization has SAML disabled.
	SamlAuthEnabled *bool `json:"saml_auth_enabled,omitempty"`
	// IdP groups determine the Konnect teams a user has.
	IdpMappingEnabled *bool `json:"idp_mapping_enabled,omitempty"`
	// A Konnect Identity Admin assigns teams to a user.
	KonnectMappingEnabled *bool `json:"konnect_mapping_enabled,omitempty"`
}

func (o *AuthenticationSettings) GetBasicAuthEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.BasicAuthEnabled
}

func (o *AuthenticationSettings) GetOidcAuthEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.OidcAuthEnabled
}

func (o *AuthenticationSettings) GetSamlAuthEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.SamlAuthEnabled
}

func (o *AuthenticationSettings) GetIdpMappingEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.IdpMappingEnabled
}

func (o *AuthenticationSettings) GetKonnectMappingEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.KonnectMappingEnabled
}
