// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteKeyAuthWithConsumerRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Consumer ID for nested entities
	ConsumerIDForNestedEntities string `pathParam:"style=simple,explode=false,name=ConsumerIdForNestedEntities"`
	// ID of the API-key to lookup
	KeyAuthID string `pathParam:"style=simple,explode=false,name=KeyAuthId"`
}

func (o *DeleteKeyAuthWithConsumerRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *DeleteKeyAuthWithConsumerRequest) GetConsumerIDForNestedEntities() string {
	if o == nil {
		return ""
	}
	return o.ConsumerIDForNestedEntities
}

func (o *DeleteKeyAuthWithConsumerRequest) GetKeyAuthID() string {
	if o == nil {
		return ""
	}
	return o.KeyAuthID
}

type DeleteKeyAuthWithConsumerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteKeyAuthWithConsumerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteKeyAuthWithConsumerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteKeyAuthWithConsumerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
