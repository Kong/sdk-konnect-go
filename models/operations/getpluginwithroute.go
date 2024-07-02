// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type GetPluginWithRouteRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Route to lookup
	RouteID string `pathParam:"style=simple,explode=false,name=RouteId"`
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
}

func (o *GetPluginWithRouteRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *GetPluginWithRouteRequest) GetRouteID() string {
	if o == nil {
		return ""
	}
	return o.RouteID
}

func (o *GetPluginWithRouteRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

type GetPluginWithRouteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched Plugin
	Plugin *components.Plugin
}

func (o *GetPluginWithRouteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPluginWithRouteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPluginWithRouteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPluginWithRouteResponse) GetPlugin() *components.Plugin {
	if o == nil {
		return nil
	}
	return o.Plugin
}