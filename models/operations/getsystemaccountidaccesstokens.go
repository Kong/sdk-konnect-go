// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var GetSystemAccountIDAccessTokensServerList = []string{
	"https://global.api.konghq.com/",
}

// GetSystemAccountIDAccessTokensQueryParamFilter - Filter access tokens returned in the response.
type GetSystemAccountIDAccessTokensQueryParamFilter struct {
	// Filter a string value field either by exact match or partial contains.
	Name *components.StringFieldFilter `queryParam:"name=name"`
}

func (o *GetSystemAccountIDAccessTokensQueryParamFilter) GetName() *components.StringFieldFilter {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetSystemAccountIDAccessTokensRequest struct {
	// ID of the system account.
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Filter access tokens returned in the response.
	Filter *GetSystemAccountIDAccessTokensQueryParamFilter `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (o *GetSystemAccountIDAccessTokensRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *GetSystemAccountIDAccessTokensRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *GetSystemAccountIDAccessTokensRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *GetSystemAccountIDAccessTokensRequest) GetFilter() *GetSystemAccountIDAccessTokensQueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type GetSystemAccountIDAccessTokensResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A paginated list response for a collection of system accounts access tokens.
	SystemAccountAccessTokenCollection *components.SystemAccountAccessTokenCollection
}

func (o *GetSystemAccountIDAccessTokensResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSystemAccountIDAccessTokensResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSystemAccountIDAccessTokensResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSystemAccountIDAccessTokensResponse) GetSystemAccountAccessTokenCollection() *components.SystemAccountAccessTokenCollection {
	if o == nil {
		return nil
	}
	return o.SystemAccountAccessTokenCollection
}
