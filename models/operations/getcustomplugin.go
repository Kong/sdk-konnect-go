// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type GetCustomPluginRequest struct {
	// ID of the CustomPlugin to lookup
	CustomPluginID string `pathParam:"style=simple,explode=false,name=CustomPluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *GetCustomPluginRequest) GetCustomPluginID() string {
	if o == nil {
		return ""
	}
	return o.CustomPluginID
}

func (o *GetCustomPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

type GetCustomPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched CustomPlugin
	CustomPlugin *components.CustomPlugin
}

func (o *GetCustomPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCustomPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCustomPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCustomPluginResponse) GetCustomPlugin() *components.CustomPlugin {
	if o == nil {
		return nil
	}
	return o.CustomPlugin
}
