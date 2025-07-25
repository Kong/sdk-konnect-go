// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/utils"
)

type CreateIdentityProviderConfigType string

const (
	CreateIdentityProviderConfigTypeConfigureOIDCIdentityProviderConfig CreateIdentityProviderConfigType = "ConfigureOIDCIdentityProviderConfig"
	CreateIdentityProviderConfigTypeSAMLIdentityProviderConfigInput     CreateIdentityProviderConfigType = "SAMLIdentityProviderConfig_input"
)

type CreateIdentityProviderConfig struct {
	ConfigureOIDCIdentityProviderConfig *ConfigureOIDCIdentityProviderConfig `queryParam:"inline"`
	SAMLIdentityProviderConfigInput     *SAMLIdentityProviderConfigInput     `queryParam:"inline"`

	Type CreateIdentityProviderConfigType
}

func CreateCreateIdentityProviderConfigConfigureOIDCIdentityProviderConfig(configureOIDCIdentityProviderConfig ConfigureOIDCIdentityProviderConfig) CreateIdentityProviderConfig {
	typ := CreateIdentityProviderConfigTypeConfigureOIDCIdentityProviderConfig

	return CreateIdentityProviderConfig{
		ConfigureOIDCIdentityProviderConfig: &configureOIDCIdentityProviderConfig,
		Type:                                typ,
	}
}

func CreateCreateIdentityProviderConfigSAMLIdentityProviderConfigInput(samlIdentityProviderConfigInput SAMLIdentityProviderConfigInput) CreateIdentityProviderConfig {
	typ := CreateIdentityProviderConfigTypeSAMLIdentityProviderConfigInput

	return CreateIdentityProviderConfig{
		SAMLIdentityProviderConfigInput: &samlIdentityProviderConfigInput,
		Type:                            typ,
	}
}

func (u *CreateIdentityProviderConfig) UnmarshalJSON(data []byte) error {

	var samlIdentityProviderConfigInput SAMLIdentityProviderConfigInput = SAMLIdentityProviderConfigInput{}
	if err := utils.UnmarshalJSON(data, &samlIdentityProviderConfigInput, "", true, true); err == nil {
		u.SAMLIdentityProviderConfigInput = &samlIdentityProviderConfigInput
		u.Type = CreateIdentityProviderConfigTypeSAMLIdentityProviderConfigInput
		return nil
	}

	var configureOIDCIdentityProviderConfig ConfigureOIDCIdentityProviderConfig = ConfigureOIDCIdentityProviderConfig{}
	if err := utils.UnmarshalJSON(data, &configureOIDCIdentityProviderConfig, "", true, true); err == nil {
		u.ConfigureOIDCIdentityProviderConfig = &configureOIDCIdentityProviderConfig
		u.Type = CreateIdentityProviderConfigTypeConfigureOIDCIdentityProviderConfig
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateIdentityProviderConfig", string(data))
}

func (u CreateIdentityProviderConfig) MarshalJSON() ([]byte, error) {
	if u.ConfigureOIDCIdentityProviderConfig != nil {
		return utils.MarshalJSON(u.ConfigureOIDCIdentityProviderConfig, "", true)
	}

	if u.SAMLIdentityProviderConfigInput != nil {
		return utils.MarshalJSON(u.SAMLIdentityProviderConfigInput, "", true)
	}

	return nil, errors.New("could not marshal union type CreateIdentityProviderConfig: all fields are null")
}

// CreateIdentityProvider - The identity provider that contains configuration data for creating an authentication integration.
type CreateIdentityProvider struct {
	// Specifies the type of identity provider.
	Type *IdentityProviderType `json:"type,omitempty"`
	// The path used for initiating login requests with the identity provider.
	LoginPath *string `json:"login_path,omitempty"`
	// Indicates whether the identity provider is enabled.
	// Only one identity provider can be active at a time, such as SAML or OIDC.
	//
	Enabled *bool                         `default:"false" json:"enabled"`
	Config  *CreateIdentityProviderConfig `json:"config,omitempty"`
}

func (c CreateIdentityProvider) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateIdentityProvider) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateIdentityProvider) GetType() *IdentityProviderType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CreateIdentityProvider) GetLoginPath() *string {
	if o == nil {
		return nil
	}
	return o.LoginPath
}

func (o *CreateIdentityProvider) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateIdentityProvider) GetConfig() *CreateIdentityProviderConfig {
	if o == nil {
		return nil
	}
	return o.Config
}
