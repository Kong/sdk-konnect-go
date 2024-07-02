// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// UpdateControlPlaneRequestAuthType - The auth type value of the cluster associated with the Runtime Group.
type UpdateControlPlaneRequestAuthType string

const (
	UpdateControlPlaneRequestAuthTypePinnedClientCerts UpdateControlPlaneRequestAuthType = "pinned_client_certs"
	UpdateControlPlaneRequestAuthTypePkiClientCerts    UpdateControlPlaneRequestAuthType = "pki_client_certs"
)

func (e UpdateControlPlaneRequestAuthType) ToPointer() *UpdateControlPlaneRequestAuthType {
	return &e
}
func (e *UpdateControlPlaneRequestAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pinned_client_certs":
		fallthrough
	case "pki_client_certs":
		*e = UpdateControlPlaneRequestAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateControlPlaneRequestAuthType: %v", v)
	}
}

// UpdateControlPlaneRequest - The request schema for the update control plane request.
type UpdateControlPlaneRequest struct {
	// The name of the control plane.
	Name *string `json:"name,omitempty"`
	// The description of the control plane in Konnect.
	Description *string `json:"description,omitempty"`
	// The auth type value of the cluster associated with the Runtime Group.
	AuthType *UpdateControlPlaneRequestAuthType `json:"auth_type,omitempty"`
	// Array of proxy URLs associated with reaching the data-planes connected to a control-plane.
	ProxyUrls []ProxyURL `json:"proxy_urls,omitempty"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	Labels map[string]string `json:"labels,omitempty"`
}

func (o *UpdateControlPlaneRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateControlPlaneRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateControlPlaneRequest) GetAuthType() *UpdateControlPlaneRequestAuthType {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *UpdateControlPlaneRequest) GetProxyUrls() []ProxyURL {
	if o == nil {
		return nil
	}
	return o.ProxyUrls
}

func (o *UpdateControlPlaneRequest) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}
