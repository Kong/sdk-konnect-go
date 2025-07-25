// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"mockserver/internal/sdk/utils"
)

// AppAuthStrategyConfigOpenIDConnect - A more advanced mode to configure an API Product Version’s Application Auth Strategy.
// Using this mode will allow developers to use API credentials issued from an external IdP that will authenticate their application requests.
// Once authenticated, an application will be granted access to any Product Version it is registered for that is configured for the same Auth Strategy.
// An OIDC strategy may be used in conjunction with a DCR provider to automatically create the IdP application.
type AppAuthStrategyConfigOpenIDConnect struct {
	Issuer               string         `json:"issuer"`
	CredentialClaim      []string       `json:"credential_claim"`
	Scopes               []string       `json:"scopes"`
	AuthMethods          []string       `json:"auth_methods"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (a AppAuthStrategyConfigOpenIDConnect) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppAuthStrategyConfigOpenIDConnect) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppAuthStrategyConfigOpenIDConnect) GetIssuer() string {
	if o == nil {
		return ""
	}
	return o.Issuer
}

func (o *AppAuthStrategyConfigOpenIDConnect) GetCredentialClaim() []string {
	if o == nil {
		return []string{}
	}
	return o.CredentialClaim
}

func (o *AppAuthStrategyConfigOpenIDConnect) GetScopes() []string {
	if o == nil {
		return []string{}
	}
	return o.Scopes
}

func (o *AppAuthStrategyConfigOpenIDConnect) GetAuthMethods() []string {
	if o == nil {
		return []string{}
	}
	return o.AuthMethods
}

func (o *AppAuthStrategyConfigOpenIDConnect) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}
