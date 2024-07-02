// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
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
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A get action response of a single assigned role.
	AssignedRole *components.AssignedRole
}

func (o *TeamsAssignRoleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TeamsAssignRoleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TeamsAssignRoleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TeamsAssignRoleResponse) GetAssignedRole() *components.AssignedRole {
	if o == nil {
		return nil
	}
	return o.AssignedRole
}
