// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type TeamsAssignRoleRequest struct {
	// The team ID
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
	// The request schema for assigning a role.
	AssignRole *components.AssignRole `request:"mediaType=application/json"`
}

func (o *TeamsAssignRoleRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *TeamsAssignRoleRequest) GetAssignRole() *components.AssignRole {
	if o == nil {
		return nil
	}
	return o.AssignRole
}

type TeamsAssignRoleResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A get action response of a single assigned role.
	AssignedRole *components.AssignedRole
}

func (o *TeamsAssignRoleResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *TeamsAssignRoleResponse) GetAssignedRole() *components.AssignedRole {
	if o == nil {
		return nil
	}
	return o.AssignedRole
}
