// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type CreateDcrProviderRequestAuth0ProviderType string

const (
	CreateDcrProviderRequestAuth0ProviderTypeAuth0 CreateDcrProviderRequestAuth0ProviderType = "auth0"
)

func (e CreateDcrProviderRequestAuth0ProviderType) ToPointer() *CreateDcrProviderRequestAuth0ProviderType {
	return &e
}
func (e *CreateDcrProviderRequestAuth0ProviderType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auth0":
		*e = CreateDcrProviderRequestAuth0ProviderType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDcrProviderRequestAuth0ProviderType: %v", v)
	}
}

// CreateDcrProviderRequestAuth0 - Request body for creating an Auth0 DCR provider.
type CreateDcrProviderRequestAuth0 struct {
	ProviderType CreateDcrProviderRequestAuth0ProviderType `json:"provider_type"`
	// Payload to create an Auth0 DCR provider.
	DcrConfig CreateDcrConfigAuth0InRequest `json:"dcr_config"`
	// The name of the DCR provider. This is used to identify the DCR provider in the Konnect UI.
	//
	Name string `json:"name"`
	// The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.
	//
	DisplayName *string `json:"display_name,omitempty"`
	Issuer      string  `json:"issuer"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	Labels map[string]string `json:"labels,omitempty"`
}

func (o *CreateDcrProviderRequestAuth0) GetProviderType() CreateDcrProviderRequestAuth0ProviderType {
	if o == nil {
		return CreateDcrProviderRequestAuth0ProviderType("")
	}
	return o.ProviderType
}

func (o *CreateDcrProviderRequestAuth0) GetDcrConfig() CreateDcrConfigAuth0InRequest {
	if o == nil {
		return CreateDcrConfigAuth0InRequest{}
	}
	return o.DcrConfig
}

func (o *CreateDcrProviderRequestAuth0) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateDcrProviderRequestAuth0) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *CreateDcrProviderRequestAuth0) GetIssuer() string {
	if o == nil {
		return ""
	}
	return o.Issuer
}

func (o *CreateDcrProviderRequestAuth0) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}
