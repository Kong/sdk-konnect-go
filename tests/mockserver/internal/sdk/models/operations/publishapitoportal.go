// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type PublishAPIToPortalRequest struct {
	// The UUID API identifier
	APIID string `pathParam:"style=simple,explode=false,name=apiId"`
	// ID of the portal.
	PortalID       string                    `pathParam:"style=simple,explode=false,name=portalId"`
	APIPublication components.APIPublication `request:"mediaType=application/json"`
}

func (o *PublishAPIToPortalRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *PublishAPIToPortalRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *PublishAPIToPortalRequest) GetAPIPublication() components.APIPublication {
	if o == nil {
		return components.APIPublication{}
	}
	return o.APIPublication
}

type PublishAPIToPortalResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// An API publication in a portal
	APIPublicationResponse *components.APIPublicationResponse
}

func (o *PublishAPIToPortalResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PublishAPIToPortalResponse) GetAPIPublicationResponse() *components.APIPublicationResponse {
	if o == nil {
		return nil
	}
	return o.APIPublicationResponse
}
