// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// IdentityProviderType - Specifies the type of identity provider.
type IdentityProviderType string

const (
	IdentityProviderTypeOidc IdentityProviderType = "oidc"
	IdentityProviderTypeSaml IdentityProviderType = "saml"
)

func (e IdentityProviderType) ToPointer() *IdentityProviderType {
	return &e
}
func (e *IdentityProviderType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oidc":
		fallthrough
	case "saml":
		*e = IdentityProviderType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IdentityProviderType: %v", v)
	}
}
