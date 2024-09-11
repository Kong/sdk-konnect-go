// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type UpsertServiceRequest struct {
	// ID of the Service to lookup
	ServiceID string `pathParam:"style=simple,explode=false,name=ServiceId"`
	// The UUID of your control plane. This variable is available in the Konnect manager
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Description of the Service
	Service components.ServiceInput `request:"mediaType=application/json"`
}

func (o *UpsertServiceRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *UpsertServiceRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpsertServiceRequest) GetService() components.ServiceInput {
	if o == nil {
		return components.ServiceInput{}
	}
	return o.Service
}

type UpsertServiceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted Service
	Service *components.Service
}

func (o *UpsertServiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertServiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertServiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertServiceResponse) GetService() *components.Service {
	if o == nil {
		return nil
	}
	return o.Service
}
