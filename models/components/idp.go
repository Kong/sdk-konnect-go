// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ClaimMappings struct {
	Name   *string `json:"name,omitempty"`
	Email  *string `json:"email,omitempty"`
	Groups *string `json:"groups,omitempty"`
}

func (o *ClaimMappings) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ClaimMappings) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *ClaimMappings) GetGroups() *string {
	if o == nil {
		return nil
	}
	return o.Groups
}

// IDP - The IdP object contains the configuration data for the OIDC authentication integration.
//
// NOTE: The `openid` scope is required. Removing it could break the OIDC integration.
type IDP struct {
	Issuer        *string        `json:"issuer,omitempty"`
	LoginPath     *string        `json:"login_path,omitempty"`
	ClientID      *string        `json:"client_id,omitempty"`
	Scopes        []string       `json:"scopes,omitempty"`
	ClaimMappings *ClaimMappings `json:"claim_mappings,omitempty"`
}

func (o *IDP) GetIssuer() *string {
	if o == nil {
		return nil
	}
	return o.Issuer
}

func (o *IDP) GetLoginPath() *string {
	if o == nil {
		return nil
	}
	return o.LoginPath
}

func (o *IDP) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *IDP) GetScopes() []string {
	if o == nil {
		return nil
	}
	return o.Scopes
}

func (o *IDP) GetClaimMappings() *ClaimMappings {
	if o == nil {
		return nil
	}
	return o.ClaimMappings
}
