// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
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
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched HMAC-auth credential
	HMACAuth *components.HMACAuth
}

func (o *GetHmacAuthResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetHmacAuthResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetHmacAuthResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetHmacAuthResponse) GetHMACAuth() *components.HMACAuth {
	if o == nil {
		return nil
	}
	return o.HMACAuth
}
