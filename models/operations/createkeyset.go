// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type CreateKeySetRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Description of the new KeySet for creation
	KeySet components.KeySetInput `request:"mediaType=application/json"`
}

func (o *CreateKeySetRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateKeySetRequest) GetKeySet() components.KeySetInput {
	if o == nil {
		return components.KeySetInput{}
	}
	return o.KeySet
}

type CreateKeySetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created KeySet
	KeySet *components.KeySet
}

func (o *CreateKeySetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateKeySetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateKeySetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateKeySetResponse) GetKeySet() *components.KeySet {
	if o == nil {
		return nil
	}
	return o.KeySet
}
