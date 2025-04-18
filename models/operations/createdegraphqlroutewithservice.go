// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type CreateDegraphqlRouteWithServiceRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Service to lookup
	ServiceID string `pathParam:"style=simple,explode=false,name=ServiceId"`
	// Description of new Degraphql_route for creation
	DegraphqlRouteWithoutParents components.DegraphqlRouteWithoutParents `request:"mediaType=application/json"`
}

func (o *CreateDegraphqlRouteWithServiceRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateDegraphqlRouteWithServiceRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateDegraphqlRouteWithServiceRequest) GetDegraphqlRouteWithoutParents() components.DegraphqlRouteWithoutParents {
	if o == nil {
		return components.DegraphqlRouteWithoutParents{}
	}
	return o.DegraphqlRouteWithoutParents
}

type CreateDegraphqlRouteWithServiceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created Degraphql_route
	DegraphqlRoute *components.DegraphqlRoute
}

func (o *CreateDegraphqlRouteWithServiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateDegraphqlRouteWithServiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateDegraphqlRouteWithServiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateDegraphqlRouteWithServiceResponse) GetDegraphqlRoute() *components.DegraphqlRoute {
	if o == nil {
		return nil
	}
	return o.DegraphqlRoute
}
