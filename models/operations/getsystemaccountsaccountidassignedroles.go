// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var GetSystemAccountsAccountIDAssignedRolesServerList = []string{
	"https://global.api.konghq.com/",
}

// GetSystemAccountsAccountIDAssignedRolesQueryParamFilter - Filter roles returned in the response.
type GetSystemAccountsAccountIDAssignedRolesQueryParamFilter struct {
	// Filter a string value by exact match.
	RoleName *string `queryParam:"name=role_name"`
	// Filter a string value by exact match.
	EntityTypeName *string `queryParam:"name=entity_type_name"`
}

func (o *GetSystemAccountsAccountIDAssignedRolesQueryParamFilter) GetRoleName() *string {
	if o == nil {
		return nil
	}
	return o.RoleName
}

func (o *GetSystemAccountsAccountIDAssignedRolesQueryParamFilter) GetEntityTypeName() *string {
	if o == nil {
		return nil
	}
	return o.EntityTypeName
}

type GetSystemAccountsAccountIDAssignedRolesRequest struct {
	// ID of the system account.
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
	// Filter roles returned in the response.
	Filter *GetSystemAccountsAccountIDAssignedRolesQueryParamFilter `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (o *GetSystemAccountsAccountIDAssignedRolesRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *GetSystemAccountsAccountIDAssignedRolesRequest) GetFilter() *GetSystemAccountsAccountIDAssignedRolesQueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type GetSystemAccountsAccountIDAssignedRolesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A paginated list response for a collection of assigned roles.
	AssignedRoleCollection *components.AssignedRoleCollection
}

func (o *GetSystemAccountsAccountIDAssignedRolesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSystemAccountsAccountIDAssignedRolesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSystemAccountsAccountIDAssignedRolesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSystemAccountsAccountIDAssignedRolesResponse) GetAssignedRoleCollection() *components.AssignedRoleCollection {
	if o == nil {
		return nil
	}
	return o.AssignedRoleCollection
}
