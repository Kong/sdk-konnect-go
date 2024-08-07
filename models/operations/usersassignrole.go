// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var UsersAssignRoleServerList = []string{
	"https://global.api.konghq.com/",
}

type UsersAssignRoleRequest struct {
	// The user ID
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
	// The request schema for assigning a role.
	AssignRole *components.AssignRole `request:"mediaType=application/json"`
}

func (o *UsersAssignRoleRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *UsersAssignRoleRequest) GetAssignRole() *components.AssignRole {
	if o == nil {
		return nil
	}
	return o.AssignRole
}

type UsersAssignRoleResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A get action response of a single assigned role.
	AssignedRole *components.AssignedRole
}

func (o *UsersAssignRoleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UsersAssignRoleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UsersAssignRoleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UsersAssignRoleResponse) GetAssignedRole() *components.AssignedRole {
	if o == nil {
		return nil
	}
	return o.AssignedRole
}
