// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var ListUserTeamsServerList = []string{
	"https://global.api.konghq.com/",
}

// ListUserTeamsQueryParamFilter - Filter teams returned in the response.
type ListUserTeamsQueryParamFilter struct {
	// Filters on the given string field value by either exact or fuzzy match.
	Name *components.StringFieldFilter `queryParam:"name=name"`
}

func (o *ListUserTeamsQueryParamFilter) GetName() *components.StringFieldFilter {
	if o == nil {
		return nil
	}
	return o.Name
}

type ListUserTeamsRequest struct {
	// The user ID.
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Filter teams returned in the response.
	Filter *ListUserTeamsQueryParamFilter `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (o *ListUserTeamsRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *ListUserTeamsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListUserTeamsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListUserTeamsRequest) GetFilter() *ListUserTeamsQueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type ListUserTeamsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A paginated list response for a collection of users.
	TeamCollection *components.TeamCollection
}

func (o *ListUserTeamsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUserTeamsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUserTeamsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListUserTeamsResponse) GetTeamCollection() *components.TeamCollection {
	if o == nil {
		return nil
	}
	return o.TeamCollection
}
