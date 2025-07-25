// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetHmacAuthRequest struct {
	// ID of the HMAC-auth credential to lookup
	HMACAuthID string `pathParam:"style=simple,explode=false,name=HMACAuthId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *GetHmacAuthRequest) GetHMACAuthID() string {
	if o == nil {
		return ""
	}
	return o.HMACAuthID
}

func (o *GetHmacAuthRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

type GetHmacAuthResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successfully fetched HMAC-auth credential
	HMACAuth *components.HMACAuth
}

func (o *GetHmacAuthResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetHmacAuthResponse) GetHMACAuth() *components.HMACAuth {
	if o == nil {
		return nil
	}
	return o.HMACAuth
}
