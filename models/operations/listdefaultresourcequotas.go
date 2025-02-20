// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var ListDefaultResourceQuotasServerList = []string{
	"https://global.api.konghq.com/",
}

type ListDefaultResourceQuotasRequest struct {
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
}

func (o *ListDefaultResourceQuotasRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListDefaultResourceQuotasRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

type ListDefaultResourceQuotasResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A paginated list for a collection of default resource quotas.
	ListDefaultResourceQuotasResponse *components.ListDefaultResourceQuotasResponse
}

func (o *ListDefaultResourceQuotasResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListDefaultResourceQuotasResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListDefaultResourceQuotasResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListDefaultResourceQuotasResponse) GetListDefaultResourceQuotasResponse() *components.ListDefaultResourceQuotasResponse {
	if o == nil {
		return nil
	}
	return o.ListDefaultResourceQuotasResponse
}
