// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type CreatePluginWithConsumerGroupRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Consumer Group to lookup
	ConsumerGroupID string `pathParam:"style=simple,explode=false,name=ConsumerGroupId"`
	// Description of new Plugin for creation
	PluginWithoutParents components.PluginWithoutParents `request:"mediaType=application/json"`
}

func (o *CreatePluginWithConsumerGroupRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreatePluginWithConsumerGroupRequest) GetConsumerGroupID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerGroupID
}

func (o *CreatePluginWithConsumerGroupRequest) GetPluginWithoutParents() components.PluginWithoutParents {
	if o == nil {
		return components.PluginWithoutParents{}
	}
	return o.PluginWithoutParents
}

type CreatePluginWithConsumerGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created Plugin
	Plugin *components.Plugin
}

func (o *CreatePluginWithConsumerGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePluginWithConsumerGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePluginWithConsumerGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreatePluginWithConsumerGroupResponse) GetPlugin() *components.Plugin {
	if o == nil {
		return nil
	}
	return o.Plugin
}
