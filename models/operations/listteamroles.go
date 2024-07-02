// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

// QueryParamFilter - Filter roles returned in the response.
type QueryParamFilter struct {
	// Filter a string value by exact match.
	RoleName *string `queryParam:"name=role_name"`
	// Filter a string value by exact match.
	EntityTypeName *string `queryParam:"name=entity_type_name"`
}

func (o *QueryParamFilter) GetRoleName() *string {
	if o == nil {
		return nil
	}
	return o.RoleName
}

func (o *QueryParamFilter) GetEntityTypeName() *string {
	if o == nil {
		return nil
	}
	return o.EntityTypeName
}

type ListTeamRolesRequest struct {
	// The team ID
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
	// Filter roles returned in the response.
	Filter *QueryParamFilter `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (o *ListTeamRolesRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *ListTeamRolesRequest) GetFilter() *QueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type ListTeamRolesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A paginated list response for a collection of assigned roles.
	AssignedRoleCollection *components.AssignedRoleCollection
}

func (o *ListTeamRolesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTeamRolesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTeamRolesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListTeamRolesResponse) GetAssignedRoleCollection() *components.AssignedRoleCollection {
	if o == nil {
		return nil
	}
	return o.AssignedRoleCollection
}