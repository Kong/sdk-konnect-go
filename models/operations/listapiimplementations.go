// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type ListAPIImplementationsRequest struct {
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Filters APIs in the response.
	Filter *components.APIImplementationFilterParameters `queryParam:"style=deepObject,explode=true,name=filter"`
	// Sorts a collection of API implementations. Supported sort attributes are:
	//
	//
	//
	//   - id
	//   - api_id
	//   - control_plane_id
	//   - service_id
	//   - created_at
	//   - updated_at
	//
	Sort *string `queryParam:"style=form,explode=true,name=sort"`
}

func (o *ListAPIImplementationsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListAPIImplementationsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListAPIImplementationsRequest) GetFilter() *components.APIImplementationFilterParameters {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *ListAPIImplementationsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

type ListAPIImplementationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Paginated list of API implementations
	ListAPIImplementationsResponse *components.ListAPIImplementationsResponse
}

func (o *ListAPIImplementationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAPIImplementationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAPIImplementationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListAPIImplementationsResponse) GetListAPIImplementationsResponse() *components.ListAPIImplementationsResponse {
	if o == nil {
		return nil
	}
	return o.ListAPIImplementationsResponse
}
