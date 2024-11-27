// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type GetKeyAuthRequest struct {
	// ID of the API-key to lookup
	KeyAuthID string `pathParam:"style=simple,explode=false,name=KeyAuthId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *GetKeyAuthRequest) GetKeyAuthID() string {
	if o == nil {
		return ""
	}
	return o.KeyAuthID
}

func (o *GetKeyAuthRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

type GetKeyAuthResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched API-key
	KeyAuth *components.KeyAuth
}

func (o *GetKeyAuthResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetKeyAuthResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetKeyAuthResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetKeyAuthResponse) GetKeyAuth() *components.KeyAuth {
	if o == nil {
		return nil
	}
	return o.KeyAuth
}