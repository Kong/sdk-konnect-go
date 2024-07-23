// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var GetSystemAccountsServerList = []string{
	"https://global.api.konghq.com/",
}

// Filter system accounts returned in the response.
type Filter struct {
	// Filter a string value field either by exact match or partial contains.
	Name *components.StringFieldFilter `queryParam:"name=name"`
	// Filter a string value field either by exact match or partial contains.
	Description *components.StringFieldFilter `queryParam:"name=description"`
	// Filter by a boolean value (true/false).
	KonnectManaged *bool `queryParam:"name=konnect_managed"`
}

func (o *Filter) GetName() *components.StringFieldFilter {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Filter) GetDescription() *components.StringFieldFilter {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Filter) GetKonnectManaged() *bool {
	if o == nil {
		return nil
	}
	return o.KonnectManaged
}

type GetSystemAccountsRequest struct {
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Filter system accounts returned in the response.
	Filter *Filter `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (o *GetSystemAccountsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *GetSystemAccountsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *GetSystemAccountsRequest) GetFilter() *Filter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type GetSystemAccountsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A paginated list response for a collection of system accounts.
	SystemAccountCollection *components.SystemAccountCollection
}

func (o *GetSystemAccountsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSystemAccountsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSystemAccountsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSystemAccountsResponse) GetSystemAccountCollection() *components.SystemAccountCollection {
	if o == nil {
		return nil
	}
	return o.SystemAccountCollection
}
