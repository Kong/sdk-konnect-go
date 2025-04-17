// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type GetDegraphqlRouteWithServiceRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Service to lookup
	ServiceID string `pathParam:"style=simple,explode=false,name=ServiceId"`
	// ID of the Degraphql_route to lookup
	DegraphqlRouteID string `pathParam:"style=simple,explode=false,name=Degraphql_routeId"`
}

func (o *GetDegraphqlRouteWithServiceRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *GetDegraphqlRouteWithServiceRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetDegraphqlRouteWithServiceRequest) GetDegraphqlRouteID() string {
	if o == nil {
		return ""
	}
	return o.DegraphqlRouteID
}

type GetDegraphqlRouteWithServiceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched Degraphql_route
	DegraphqlRoute *components.DegraphqlRoute
}

func (o *GetDegraphqlRouteWithServiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDegraphqlRouteWithServiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDegraphqlRouteWithServiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetDegraphqlRouteWithServiceResponse) GetDegraphqlRoute() *components.DegraphqlRoute {
	if o == nil {
		return nil
	}
	return o.DegraphqlRoute
}
