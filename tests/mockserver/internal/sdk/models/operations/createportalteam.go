// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type CreatePortalTeamRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
	// Create a team in a portal.
	PortalCreateTeamRequest *components.PortalCreateTeamRequest `request:"mediaType=application/json"`
}

func (o *CreatePortalTeamRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *CreatePortalTeamRequest) GetPortalCreateTeamRequest() *components.PortalCreateTeamRequest {
	if o == nil {
		return nil
	}
	return o.PortalCreateTeamRequest
}

type CreatePortalTeamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Details about a team of developers in a portal.
	PortalTeamResponse *components.PortalTeamResponse
}

func (o *CreatePortalTeamResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreatePortalTeamResponse) GetPortalTeamResponse() *components.PortalTeamResponse {
	if o == nil {
		return nil
	}
	return o.PortalTeamResponse
}
