// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/utils"
)

type SchemasCreateIdentityProviderConfigType string

const (
	SchemasCreateIdentityProviderConfigTypeConfigureOIDCIdentityProviderConfig    SchemasCreateIdentityProviderConfigType = "ConfigureOIDCIdentityProviderConfig"
	SchemasCreateIdentityProviderConfigTypeSchemasSAMLIdentityProviderConfigInput SchemasCreateIdentityProviderConfigType = "schemas-SAMLIdentityProviderConfig_input"
)

type SchemasCreateIdentityProviderConfig struct {
	ConfigureOIDCIdentityProviderConfig    *ConfigureOIDCIdentityProviderConfig    `queryParam:"inline"`
	SchemasSAMLIdentityProviderConfigInput *SchemasSAMLIdentityProviderConfigInput `queryParam:"inline"`

	Type SchemasCreateIdentityProviderConfigType
}

func CreateSchemasCreateIdentityProviderConfigConfigureOIDCIdentityProviderConfig(configureOIDCIdentityProviderConfig ConfigureOIDCIdentityProviderConfig) SchemasCreateIdentityProviderConfig {
	typ := SchemasCreateIdentityProviderConfigTypeConfigureOIDCIdentityProviderConfig

	return SchemasCreateIdentityProviderConfig{
		ConfigureOIDCIdentityProviderConfig: &configureOIDCIdentityProviderConfig,
		Type:                                typ,
	}
}

func CreateSchemasCreateIdentityProviderConfigSchemasSAMLIdentityProviderConfigInput(schemasSAMLIdentityProviderConfigInput SchemasSAMLIdentityProviderConfigInput) SchemasCreateIdentityProviderConfig {
	typ := SchemasCreateIdentityProviderConfigTypeSchemasSAMLIdentityProviderConfigInput

	return SchemasCreateIdentityProviderConfig{
		SchemasSAMLIdentityProviderConfigInput: &schemasSAMLIdentityProviderConfigInput,
		Type:                                   typ,
	}
}

func (u *SchemasCreateIdentityProviderConfig) UnmarshalJSON(data []byte) error {

	var schemasSAMLIdentityProviderConfigInput SchemasSAMLIdentityProviderConfigInput = SchemasSAMLIdentityProviderConfigInput{}
	if err := utils.UnmarshalJSON(data, &schemasSAMLIdentityProviderConfigInput, "", true, true); err == nil {
		u.SchemasSAMLIdentityProviderConfigInput = &schemasSAMLIdentityProviderConfigInput
		u.Type = SchemasCreateIdentityProviderConfigTypeSchemasSAMLIdentityProviderConfigInput
		return nil
	}

	var configureOIDCIdentityProviderConfig ConfigureOIDCIdentityProviderConfig = ConfigureOIDCIdentityProviderConfig{}
	if err := utils.UnmarshalJSON(data, &configureOIDCIdentityProviderConfig, "", true, true); err == nil {
		u.ConfigureOIDCIdentityProviderConfig = &configureOIDCIdentityProviderConfig
		u.Type = SchemasCreateIdentityProviderConfigTypeConfigureOIDCIdentityProviderConfig
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SchemasCreateIdentityProviderConfig", string(data))
}

func (u SchemasCreateIdentityProviderConfig) MarshalJSON() ([]byte, error) {
	if u.ConfigureOIDCIdentityProviderConfig != nil {
		return utils.MarshalJSON(u.ConfigureOIDCIdentityProviderConfig, "", true)
	}

	if u.SchemasSAMLIdentityProviderConfigInput != nil {
		return utils.MarshalJSON(u.SchemasSAMLIdentityProviderConfigInput, "", true)
	}

	return nil, errors.New("could not marshal union type SchemasCreateIdentityProviderConfig: all fields are null")
}

// SchemasCreateIdentityProvider - The identity provider that contains configuration data for creating an authentication integration.
type SchemasCreateIdentityProvider struct {
	// Indicates whether the identity provider is enabled.
	// Only one identity provider can be active at a time, such as SAML or OIDC.
	//
	Enabled *bool `default:"false" json:"enabled"`
	// Specifies the type of identity provider.
	Type   *IdentityProviderType                `json:"type,omitempty"`
	Config *SchemasCreateIdentityProviderConfig `json:"config,omitempty"`
}

func (s SchemasCreateIdentityProvider) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SchemasCreateIdentityProvider) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SchemasCreateIdentityProvider) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *SchemasCreateIdentityProvider) GetType() *IdentityProviderType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SchemasCreateIdentityProvider) GetConfig() *SchemasCreateIdentityProviderConfig {
	if o == nil {
		return nil
	}
	return o.Config
}
