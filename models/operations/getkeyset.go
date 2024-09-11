// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type GetKeySetRequest struct {
	// ID of the KeySet to lookup
	KeySetID string `pathParam:"style=simple,explode=false,name=KeySetId"`
	// The UUID of your control plane. This variable is available in the Konnect manager
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *GetKeySetRequest) GetKeySetID() string {
	if o == nil {
		return ""
	}
	return o.KeySetID
}

func (o *GetKeySetRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

type GetKeySetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched KeySet
	KeySet *components.KeySet
}

func (o *GetKeySetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetKeySetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetKeySetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetKeySetResponse) GetKeySet() *components.KeySet {
	if o == nil {
		return nil
	}
	return o.KeySet
}
