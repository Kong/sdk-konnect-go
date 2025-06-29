// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type CreatePluginWithConsumerRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Consumer ID for nested entities
	ConsumerIDForNestedEntities string `pathParam:"style=simple,explode=false,name=ConsumerIdForNestedEntities"`
	// Description of new Plugin for creation
	PluginWithoutParents components.PluginWithoutParents `request:"mediaType=application/json"`
}

func (o *CreatePluginWithConsumerRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreatePluginWithConsumerRequest) GetConsumerIDForNestedEntities() string {
	if o == nil {
		return ""
	}
	return o.ConsumerIDForNestedEntities
}

func (o *CreatePluginWithConsumerRequest) GetPluginWithoutParents() components.PluginWithoutParents {
	if o == nil {
		return components.PluginWithoutParents{}
	}
	return o.PluginWithoutParents
}

type CreatePluginWithConsumerResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successfully created Plugin
	Plugin *components.Plugin
}

func (o *CreatePluginWithConsumerResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreatePluginWithConsumerResponse) GetPlugin() *components.Plugin {
	if o == nil {
		return nil
	}
	return o.Plugin
}
