// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type CreatePluginSchemasRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID      string                          `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreatePluginSchemas *components.CreatePluginSchemas `request:"mediaType=application/json"`
}

func (o *CreatePluginSchemasRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreatePluginSchemasRequest) GetCreatePluginSchemas() *components.CreatePluginSchemas {
	if o == nil {
		return nil
	}
	return o.CreatePluginSchemas
}

type CreatePluginSchemasResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A response for a single custom plugin schema.
	PluginSchemas *components.PluginSchemas
}

func (o *CreatePluginSchemasResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreatePluginSchemasResponse) GetPluginSchemas() *components.PluginSchemas {
	if o == nil {
		return nil
	}
	return o.PluginSchemas
}
