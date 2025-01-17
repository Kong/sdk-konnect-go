// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
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
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched MTLS-auth credential
	MTLSAuth *components.MTLSAuth
}

func (o *GetMtlsAuthResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMtlsAuthResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMtlsAuthResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMtlsAuthResponse) GetMTLSAuth() *components.MTLSAuth {
	if o == nil {
		return nil
	}
	return o.MTLSAuth
}
