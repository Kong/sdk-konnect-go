// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetControlPlanesIDGroupMemberStatusRequest struct {
	// ID of a control plane
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetControlPlanesIDGroupMemberStatusRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetControlPlanesIDGroupMemberStatusResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Determines the group membership status of a control plane.
	GroupMemberStatus *components.GroupMemberStatus
}

func (o *GetControlPlanesIDGroupMemberStatusResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetControlPlanesIDGroupMemberStatusResponse) GetGroupMemberStatus() *components.GroupMemberStatus {
	if o == nil {
		return nil
	}
	return o.GroupMemberStatus
}
