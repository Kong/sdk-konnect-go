// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type CreatePluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Description of the new Plugin for creation
	Plugin components.PluginInput `request:"mediaType=application/json"`
}

func (o *CreatePluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreatePluginRequest) GetPlugin() components.PluginInput {
	if o == nil {
		return components.PluginInput{}
	}
	return o.Plugin
}

type CreatePluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created Plugin
	Plugin *components.Plugin
}

func (o *CreatePluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreatePluginResponse) GetPlugin() *components.Plugin {
	if o == nil {
		return nil
	}
	return o.Plugin
}
