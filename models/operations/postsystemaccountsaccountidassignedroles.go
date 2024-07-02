// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type PostSystemAccountsAccountIDAssignedRolesRequest struct {
	// ID of the system account.
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
	// The request schema for assigning a role.
	AssignRole *components.AssignRole `request:"mediaType=application/json"`
}

func (o *PostSystemAccountsAccountIDAssignedRolesRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *PostSystemAccountsAccountIDAssignedRolesRequest) GetAssignRole() *components.AssignRole {
	if o == nil {
		return nil
	}
	return o.AssignRole
}

type PostSystemAccountsAccountIDAssignedRolesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A get action response of a single assigned role.
	AssignedRole *components.AssignedRole
}

func (o *PostSystemAccountsAccountIDAssignedRolesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostSystemAccountsAccountIDAssignedRolesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostSystemAccountsAccountIDAssignedRolesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostSystemAccountsAccountIDAssignedRolesResponse) GetAssignedRole() *components.AssignedRole {
	if o == nil {
		return nil
	}
	return o.AssignedRole
}