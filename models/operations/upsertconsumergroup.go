// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type UpsertConsumerGroupRequest struct {
	// ID of the Consumer Group to lookup
	ConsumerGroupID string `pathParam:"style=simple,explode=false,name=ConsumerGroupId"`
	// The UUID of your control plane. This variable is available in the Konnect manager
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Description of the Consumer Group
	ConsumerGroup components.ConsumerGroupInput `request:"mediaType=application/json"`
}

func (o *UpsertConsumerGroupRequest) GetConsumerGroupID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerGroupID
}

func (o *UpsertConsumerGroupRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpsertConsumerGroupRequest) GetConsumerGroup() components.ConsumerGroupInput {
	if o == nil {
		return components.ConsumerGroupInput{}
	}
	return o.ConsumerGroup
}

type UpsertConsumerGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted Consumer Group
	ConsumerGroup *components.ConsumerGroup
}

func (o *UpsertConsumerGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertConsumerGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertConsumerGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertConsumerGroupResponse) GetConsumerGroup() *components.ConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}
