// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type ListAPIVersionsRequest struct {
	// The UUID API identifier
	APIID string `pathParam:"style=simple,explode=false,name=apiId"`
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Filters API Version in the response.
	Filter *components.APIVersionFilterParameters `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (o *ListAPIVersionsRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *ListAPIVersionsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListAPIVersionsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListAPIVersionsRequest) GetFilter() *components.APIVersionFilterParameters {
	if o == nil {
		return nil
	}
	return o.Filter
}

type ListAPIVersionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List of API specifications (OpenAPI or AsyncAPI)
	ListAPIVersionResponse *components.ListAPIVersionResponse
}

func (o *ListAPIVersionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAPIVersionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAPIVersionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListAPIVersionsResponse) GetListAPIVersionResponse() *components.ListAPIVersionResponse {
	if o == nil {
		return nil
	}
	return o.ListAPIVersionResponse
}
