// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"time"
)

type IdentityProviderConfigType string

const (
	IdentityProviderConfigTypeOIDCIdentityProviderConfig IdentityProviderConfigType = "OIDCIdentityProviderConfig"
	IdentityProviderConfigTypeSAMLIdentityProviderConfig IdentityProviderConfigType = "SAMLIdentityProviderConfig"
)

type IdentityProviderConfig struct {
	OIDCIdentityProviderConfig *OIDCIdentityProviderConfig `queryParam:"inline"`
	SAMLIdentityProviderConfig *SAMLIdentityProviderConfig `queryParam:"inline"`

	Type IdentityProviderConfigType
}

func CreateIdentityProviderConfigOIDCIdentityProviderConfig(oidcIdentityProviderConfig OIDCIdentityProviderConfig) IdentityProviderConfig {
	typ := IdentityProviderConfigTypeOIDCIdentityProviderConfig

	return IdentityProviderConfig{
		OIDCIdentityProviderConfig: &oidcIdentityProviderConfig,
		Type:                       typ,
	}
}

func CreateIdentityProviderConfigSAMLIdentityProviderConfig(samlIdentityProviderConfig SAMLIdentityProviderConfig) IdentityProviderConfig {
	typ := IdentityProviderConfigTypeSAMLIdentityProviderConfig

	return IdentityProviderConfig{
		SAMLIdentityProviderConfig: &samlIdentityProviderConfig,
		Type:                       typ,
	}
}

func (u *IdentityProviderConfig) UnmarshalJSON(data []byte) error {

	var oidcIdentityProviderConfig OIDCIdentityProviderConfig = OIDCIdentityProviderConfig{}
	if err := utils.UnmarshalJSON(data, &oidcIdentityProviderConfig, "", true, true); err == nil {
		u.OIDCIdentityProviderConfig = &oidcIdentityProviderConfig
		u.Type = IdentityProviderConfigTypeOIDCIdentityProviderConfig
		return nil
	}

	var samlIdentityProviderConfig SAMLIdentityProviderConfig = SAMLIdentityProviderConfig{}
	if err := utils.UnmarshalJSON(data, &samlIdentityProviderConfig, "", true, true); err == nil {
		u.SAMLIdentityProviderConfig = &samlIdentityProviderConfig
		u.Type = IdentityProviderConfigTypeSAMLIdentityProviderConfig
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for IdentityProviderConfig", string(data))
}

func (u IdentityProviderConfig) MarshalJSON() ([]byte, error) {
	if u.OIDCIdentityProviderConfig != nil {
		return utils.MarshalJSON(u.OIDCIdentityProviderConfig, "", true)
	}

	if u.SAMLIdentityProviderConfig != nil {
		return utils.MarshalJSON(u.SAMLIdentityProviderConfig, "", true)
	}

	return nil, errors.New("could not marshal union type IdentityProviderConfig: all fields are null")
}

// IdentityProvider - The identity provider that contains configuration data for authentication integration.
type IdentityProvider struct {
	// Contains a unique identifier used for this resource.
	ID *string `json:"id,omitempty"`
	// Specifies the type of identity provider.
	Type *IdentityProviderType `json:"type,omitempty"`
	// Indicates whether the identity provider is enabled.
	// Only one identity provider can be active at a time, such as SAML or OIDC.
	//
	Enabled *bool `default:"false" json:"enabled"`
	// The path used for initiating login requests with the identity provider.
	LoginPath *string                 `json:"login_path,omitempty"`
	Config    *IdentityProviderConfig `json:"config,omitempty"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (i IdentityProvider) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IdentityProvider) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IdentityProvider) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IdentityProvider) GetType() *IdentityProviderType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *IdentityProvider) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *IdentityProvider) GetLoginPath() *string {
	if o == nil {
		return nil
	}
	return o.LoginPath
}

func (o *IdentityProvider) GetConfig() *IdentityProviderConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *IdentityProvider) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *IdentityProvider) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
