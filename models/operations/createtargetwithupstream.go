// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type CreateTargetWithUpstreamRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID or target of the Target to lookup
	UpstreamIDForTarget string `pathParam:"style=simple,explode=false,name=UpstreamIdForTarget"`
	// Description of new Target for creation
	TargetWithoutParents components.TargetWithoutParents `request:"mediaType=application/json"`
}

func (o *CreateTargetWithUpstreamRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateTargetWithUpstreamRequest) GetUpstreamIDForTarget() string {
	if o == nil {
		return ""
	}
	return o.UpstreamIDForTarget
}

func (o *CreateTargetWithUpstreamRequest) GetTargetWithoutParents() components.TargetWithoutParents {
	if o == nil {
		return components.TargetWithoutParents{}
	}
	return o.TargetWithoutParents
}

type CreateTargetWithUpstreamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created Target
	Target *components.Target
}

func (o *CreateTargetWithUpstreamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTargetWithUpstreamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTargetWithUpstreamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTargetWithUpstreamResponse) GetTarget() *components.Target {
	if o == nil {
		return nil
	}
	return o.Target
}
