// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type UpdateDeveloperRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
	// ID of the developer.
	DeveloperID string `pathParam:"style=simple,explode=false,name=developerId"`
	// Update a developer.
	UpdateDeveloperRequest components.UpdateDeveloperRequest `request:"mediaType=application/json"`
}

func (o *UpdateDeveloperRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *UpdateDeveloperRequest) GetDeveloperID() string {
	if o == nil {
		return ""
	}
	return o.DeveloperID
}

func (o *UpdateDeveloperRequest) GetUpdateDeveloperRequest() components.UpdateDeveloperRequest {
	if o == nil {
		return components.UpdateDeveloperRequest{}
	}
	return o.UpdateDeveloperRequest
}

type UpdateDeveloperResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Details about the developer being updated.
	PortalDeveloper *components.PortalDeveloper
}

func (o *UpdateDeveloperResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateDeveloperResponse) GetPortalDeveloper() *components.PortalDeveloper {
	if o == nil {
		return nil
	}
	return o.PortalDeveloper
}
