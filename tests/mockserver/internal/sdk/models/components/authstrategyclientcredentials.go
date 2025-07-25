// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/utils"
)

type AuthStrategyClientCredentialsCredentialType string

const (
	AuthStrategyClientCredentialsCredentialTypeClientCredentials            AuthStrategyClientCredentialsCredentialType = "client_credentials"
	AuthStrategyClientCredentialsCredentialTypeSelfManagedClientCredentials AuthStrategyClientCredentialsCredentialType = "self_managed_client_credentials"
)

func (e AuthStrategyClientCredentialsCredentialType) ToPointer() *AuthStrategyClientCredentialsCredentialType {
	return &e
}
func (e *AuthStrategyClientCredentialsCredentialType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "client_credentials":
		fallthrough
	case "self_managed_client_credentials":
		*e = AuthStrategyClientCredentialsCredentialType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthStrategyClientCredentialsCredentialType: %v", v)
	}
}

// AuthStrategyClientCredentials - Client Credential Auth strategy that the application uses.
type AuthStrategyClientCredentials struct {
	// The Application Auth Strategy ID.
	ID             string                                      `json:"id"`
	Name           *string                                     `default:"name" json:"name"`
	CredentialType AuthStrategyClientCredentialsCredentialType `json:"credential_type"`
	AuthMethods    []string                                    `json:"auth_methods"`
	// Possible developer selectable scopes for an application. Only present when using DCR Provider that supports it.
	AvailableScopes []string `json:"available_scopes,omitempty"`
}

func (a AuthStrategyClientCredentials) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthStrategyClientCredentials) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AuthStrategyClientCredentials) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AuthStrategyClientCredentials) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AuthStrategyClientCredentials) GetCredentialType() AuthStrategyClientCredentialsCredentialType {
	if o == nil {
		return AuthStrategyClientCredentialsCredentialType("")
	}
	return o.CredentialType
}

func (o *AuthStrategyClientCredentials) GetAuthMethods() []string {
	if o == nil {
		return []string{}
	}
	return o.AuthMethods
}

func (o *AuthStrategyClientCredentials) GetAvailableScopes() []string {
	if o == nil {
		return nil
	}
	return o.AvailableScopes
}
