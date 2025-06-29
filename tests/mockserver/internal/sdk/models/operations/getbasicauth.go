// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetBasicAuthRequest struct {
	// ID of the Basic-auth credential to lookup
	BasicAuthID string `pathParam:"style=simple,explode=false,name=BasicAuthId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *GetBasicAuthRequest) GetBasicAuthID() string {
	if o == nil {
		return ""
	}
	return o.BasicAuthID
}

func (o *GetBasicAuthRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

type GetBasicAuthResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successfully fetched Basic-auth credential
	BasicAuth *components.BasicAuth
}

func (o *GetBasicAuthResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetBasicAuthResponse) GetBasicAuth() *components.BasicAuth {
	if o == nil {
		return nil
	}
	return o.BasicAuth
}
