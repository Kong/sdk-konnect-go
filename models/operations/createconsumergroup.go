// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type CreateConsumerGroupRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Description of the new Consumer Group for creation
	ConsumerGroup components.ConsumerGroup `request:"mediaType=application/json"`
}

func (o *CreateConsumerGroupRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateConsumerGroupRequest) GetConsumerGroup() components.ConsumerGroup {
	if o == nil {
		return components.ConsumerGroup{}
	}
	return o.ConsumerGroup
}

type CreateConsumerGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created Consumer Group
	ConsumerGroup *components.ConsumerGroup
}

func (o *CreateConsumerGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateConsumerGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateConsumerGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateConsumerGroupResponse) GetConsumerGroup() *components.ConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}
