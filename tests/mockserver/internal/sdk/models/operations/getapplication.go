// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetApplicationRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
	// ID of the application.
	ApplicationID string `pathParam:"style=simple,explode=false,name=applicationId"`
}

func (o *GetApplicationRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *GetApplicationRequest) GetApplicationID() string {
	if o == nil {
		return ""
	}
	return o.ApplicationID
}

type GetApplicationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Details about an application in a portal.
	GetApplicationResponse *components.GetApplicationResponse
}

func (o *GetApplicationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetApplicationResponse) GetGetApplicationResponse() *components.GetApplicationResponse {
	if o == nil {
		return nil
	}
	return o.GetApplicationResponse
}
