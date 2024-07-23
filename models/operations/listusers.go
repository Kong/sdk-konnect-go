// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var ListUsersServerList = []string{
	"https://global.api.konghq.com/",
}

// ListUsersQueryParamFilter - Filter users returned in the response.
type ListUsersQueryParamFilter struct {
	// Filter a string value by exact match.
	ID *string `queryParam:"name=id"`
	// Filter a string value field either by exact match or partial contains.
	Email *components.StringFieldFilter `queryParam:"name=email"`
	// Filter a string value field either by exact match or partial contains.
	FullName *components.StringFieldFilter `queryParam:"name=full_name"`
	// Filter by a boolean value (true/false).
	Active *bool `queryParam:"name=active"`
}

func (o *ListUsersQueryParamFilter) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ListUsersQueryParamFilter) GetEmail() *components.StringFieldFilter {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *ListUsersQueryParamFilter) GetFullName() *components.StringFieldFilter {
	if o == nil {
		return nil
	}
	return o.FullName
}

func (o *ListUsersQueryParamFilter) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

type ListUsersRequest struct {
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Filter users returned in the response.
	Filter *ListUsersQueryParamFilter `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (o *ListUsersRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListUsersRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListUsersRequest) GetFilter() *ListUsersQueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type ListUsersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A paginated list response for a collection of users.
	UserCollection *components.UserCollection
}

func (o *ListUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUsersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUsersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListUsersResponse) GetUserCollection() *components.UserCollection {
	if o == nil {
		return nil
	}
	return o.UserCollection
}
