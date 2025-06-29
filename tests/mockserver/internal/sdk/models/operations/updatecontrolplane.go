// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type UpdateControlPlaneRequest struct {
	// The control plane ID
	ID                        string                               `pathParam:"style=simple,explode=false,name=id"`
	UpdateControlPlaneRequest components.UpdateControlPlaneRequest `request:"mediaType=application/json"`
}

func (o *UpdateControlPlaneRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateControlPlaneRequest) GetUpdateControlPlaneRequest() components.UpdateControlPlaneRequest {
	if o == nil {
		return components.UpdateControlPlaneRequest{}
	}
	return o.UpdateControlPlaneRequest
}

type UpdateControlPlaneResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A response to updating a control plane.
	ControlPlane *components.ControlPlane
}

func (o *UpdateControlPlaneResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateControlPlaneResponse) GetControlPlane() *components.ControlPlane {
	if o == nil {
		return nil
	}
	return o.ControlPlane
}
