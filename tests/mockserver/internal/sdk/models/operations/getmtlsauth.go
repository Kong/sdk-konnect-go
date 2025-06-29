// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetMtlsAuthRequest struct {
	// ID of the MTLS-auth credential to lookup
	MTLSAuthID string `pathParam:"style=simple,explode=false,name=MTLSAuthId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *GetMtlsAuthRequest) GetMTLSAuthID() string {
	if o == nil {
		return ""
	}
	return o.MTLSAuthID
}

func (o *GetMtlsAuthRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

type GetMtlsAuthResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successfully fetched MTLS-auth credential
	MTLSAuth *components.MTLSAuth
}

func (o *GetMtlsAuthResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMtlsAuthResponse) GetMTLSAuth() *components.MTLSAuth {
	if o == nil {
		return nil
	}
	return o.MTLSAuth
}
