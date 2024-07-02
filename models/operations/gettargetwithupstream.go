// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type GetTargetWithUpstreamRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID or target of the Target to lookup
	UpstreamIDForTarget string `pathParam:"style=simple,explode=false,name=UpstreamIdForTarget"`
	// ID of the Target to lookup
	TargetID string `pathParam:"style=simple,explode=false,name=TargetId"`
}

func (o *GetTargetWithUpstreamRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *GetTargetWithUpstreamRequest) GetUpstreamIDForTarget() string {
	if o == nil {
		return ""
	}
	return o.UpstreamIDForTarget
}

func (o *GetTargetWithUpstreamRequest) GetTargetID() string {
	if o == nil {
		return ""
	}
	return o.TargetID
}

type GetTargetWithUpstreamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched Target
	Target *components.Target
}

func (o *GetTargetWithUpstreamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTargetWithUpstreamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTargetWithUpstreamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTargetWithUpstreamResponse) GetTarget() *components.Target {
	if o == nil {
		return nil
	}
	return o.Target
}