// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type UpdatePortalPageRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
	// ID of the page.
	PageID string `pathParam:"style=simple,explode=false,name=pageId"`
	// Update a page in a portal.
	UpdatePortalPageRequest components.UpdatePortalPageRequest `request:"mediaType=application/json"`
}

func (o *UpdatePortalPageRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *UpdatePortalPageRequest) GetPageID() string {
	if o == nil {
		return ""
	}
	return o.PageID
}

func (o *UpdatePortalPageRequest) GetUpdatePortalPageRequest() components.UpdatePortalPageRequest {
	if o == nil {
		return components.UpdatePortalPageRequest{}
	}
	return o.UpdatePortalPageRequest
}

type UpdatePortalPageResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Details about a page in a portal.
	PortalPageResponse *components.PortalPageResponse
}

func (o *UpdatePortalPageResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdatePortalPageResponse) GetPortalPageResponse() *components.PortalPageResponse {
	if o == nil {
		return nil
	}
	return o.PortalPageResponse
}
