// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveConsumerFromConsumerGroupRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Consumer to lookup
	ConsumerID string `pathParam:"style=simple,explode=false,name=ConsumerId"`
	// ID of the Consumer Group to lookup
	ConsumerGroupID string `pathParam:"style=simple,explode=false,name=ConsumerGroupId"`
}

func (o *RemoveConsumerFromConsumerGroupRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *RemoveConsumerFromConsumerGroupRequest) GetConsumerID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerID
}

func (o *RemoveConsumerFromConsumerGroupRequest) GetConsumerGroupID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerGroupID
}

type RemoveConsumerFromConsumerGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RemoveConsumerFromConsumerGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveConsumerFromConsumerGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveConsumerFromConsumerGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
