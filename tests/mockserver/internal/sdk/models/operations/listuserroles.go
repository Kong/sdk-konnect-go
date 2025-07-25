// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

// ListUserRolesFilter - Filter roles returned in the response.
type ListUserRolesFilter struct {
	// Filters on the given string field value by exact match.
	RoleName *components.StringFieldEqualsFilterUnion `queryParam:"name=role_name"`
	// Filters on the given string field value by exact match.
	EntityTypeName *components.StringFieldEqualsFilterUnion `queryParam:"name=entity_type_name"`
}

func (o *ListUserRolesFilter) GetRoleName() *components.StringFieldEqualsFilterUnion {
	if o == nil {
		return nil
	}
	return o.RoleName
}

func (o *ListUserRolesFilter) GetEntityTypeName() *components.StringFieldEqualsFilterUnion {
	if o == nil {
		return nil
	}
	return o.EntityTypeName
}

type ListUserRolesRequest struct {
	// The user ID
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
	// Filter roles returned in the response.
	Filter *ListUserRolesFilter `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (o *ListUserRolesRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *ListUserRolesRequest) GetFilter() *ListUserRolesFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type ListUserRolesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A paginated list response for a collection of assigned roles.
	AssignedRoleCollection *components.AssignedRoleCollection
}

func (o *ListUserRolesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListUserRolesResponse) GetAssignedRoleCollection() *components.AssignedRoleCollection {
	if o == nil {
		return nil
	}
	return o.AssignedRoleCollection
}
