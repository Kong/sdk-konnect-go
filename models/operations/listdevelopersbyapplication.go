// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

// ListDevelopersByApplicationQueryParamFilter - Filter application developers returned in the response.
type ListDevelopersByApplicationQueryParamFilter struct {
	// Filters on the given string field value by either exact or fuzzy match.
	ID *components.StringFieldFilter `queryParam:"name=id"`
}

func (o *ListDevelopersByApplicationQueryParamFilter) GetID() *components.StringFieldFilter {
	if o == nil {
		return nil
	}
	return o.ID
}

type ListDevelopersByApplicationRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
	// ID of the application.
	ApplicationID string `pathParam:"style=simple,explode=false,name=applicationId"`
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Sorts a set of developers for an application. Supported sort attributes are:
	//
	//
	//
	//   - id
	//
	Sort *string `queryParam:"style=form,explode=true,name=sort"`
	// Filter application developers returned in the response.
	Filter *ListDevelopersByApplicationQueryParamFilter `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (o *ListDevelopersByApplicationRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *ListDevelopersByApplicationRequest) GetApplicationID() string {
	if o == nil {
		return ""
	}
	return o.ApplicationID
}

func (o *ListDevelopersByApplicationRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListDevelopersByApplicationRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListDevelopersByApplicationRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListDevelopersByApplicationRequest) GetFilter() *ListDevelopersByApplicationQueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type ListDevelopersByApplicationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A paginated list developers of an application.
	ListApplicationDevelopersResponse *components.ListApplicationDevelopersResponse
}

func (o *ListDevelopersByApplicationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListDevelopersByApplicationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListDevelopersByApplicationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListDevelopersByApplicationResponse) GetListApplicationDevelopersResponse() *components.ListApplicationDevelopersResponse {
	if o == nil {
		return nil
	}
	return o.ListApplicationDevelopersResponse
}
