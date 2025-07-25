// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type UpdatePortalAuthenticationSettingsRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
	// Update a portal's developer authentication settings.
	PortalAuthenticationSettingsUpdateRequest *components.PortalAuthenticationSettingsUpdateRequest `request:"mediaType=application/json"`
}

func (o *UpdatePortalAuthenticationSettingsRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *UpdatePortalAuthenticationSettingsRequest) GetPortalAuthenticationSettingsUpdateRequest() *components.PortalAuthenticationSettingsUpdateRequest {
	if o == nil {
		return nil
	}
	return o.PortalAuthenticationSettingsUpdateRequest
}

type UpdatePortalAuthenticationSettingsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Details about a portal's authentication settings.
	PortalAuthenticationSettingsResponse *components.PortalAuthenticationSettingsResponse
}

func (o *UpdatePortalAuthenticationSettingsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdatePortalAuthenticationSettingsResponse) GetPortalAuthenticationSettingsResponse() *components.PortalAuthenticationSettingsResponse {
	if o == nil {
		return nil
	}
	return o.PortalAuthenticationSettingsResponse
}
