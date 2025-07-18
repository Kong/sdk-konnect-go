// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetPortalRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
}

func (o *GetPortalRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

type GetPortalResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Details about a portal.
	PortalResponse *components.PortalResponseResponse
}

func (o *GetPortalResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPortalResponse) GetPortalResponse() *components.PortalResponseResponse {
	if o == nil {
		return nil
	}
	return o.PortalResponse
}
