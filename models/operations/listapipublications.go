// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type ListAPIPublicationsRequest struct {
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Filters API Publications in the response.
	Filter *components.APIPublicationFilterParameters `queryParam:"style=deepObject,explode=true,name=filter"`
	// Sorts a collection of API publications. Supported sort attributes are:
	//   - portal_id
	//   - portal_name
	//   - api_id
	//   - api_name
	//   - created_at
	//   - updated_at
	//
	Sort *string `queryParam:"style=form,explode=true,name=sort"`
}

func (o *ListAPIPublicationsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListAPIPublicationsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListAPIPublicationsRequest) GetFilter() *components.APIPublicationFilterParameters {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *ListAPIPublicationsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

type ListAPIPublicationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Paginated list of API publications
	ListAPIPublicationResponse *components.ListAPIPublicationResponse
}

func (o *ListAPIPublicationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAPIPublicationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAPIPublicationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListAPIPublicationsResponse) GetListAPIPublicationResponse() *components.ListAPIPublicationResponse {
	if o == nil {
		return nil
	}
	return o.ListAPIPublicationResponse
}
