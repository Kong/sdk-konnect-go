// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type UpsertKeyWithKeySetRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the KeySet to lookup
	KeySetID string `pathParam:"style=simple,explode=false,name=KeySetId"`
	// ID of the Key to lookup
	KeyID string `pathParam:"style=simple,explode=false,name=KeyId"`
	// Description of the Key
	KeyWithoutParents components.KeyWithoutParents `request:"mediaType=application/json"`
}

func (o *UpsertKeyWithKeySetRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpsertKeyWithKeySetRequest) GetKeySetID() string {
	if o == nil {
		return ""
	}
	return o.KeySetID
}

func (o *UpsertKeyWithKeySetRequest) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *UpsertKeyWithKeySetRequest) GetKeyWithoutParents() components.KeyWithoutParents {
	if o == nil {
		return components.KeyWithoutParents{}
	}
	return o.KeyWithoutParents
}

type UpsertKeyWithKeySetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted Key
	Key *components.Key
}

func (o *UpsertKeyWithKeySetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertKeyWithKeySetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertKeyWithKeySetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertKeyWithKeySetResponse) GetKey() *components.Key {
	if o == nil {
		return nil
	}
	return o.Key
}