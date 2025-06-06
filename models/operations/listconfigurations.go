// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var ListConfigurationsServerList = []string{
	"https://global.api.konghq.com/",
}

type ListConfigurationsRequest struct {
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Filters supported for configurations.
	Filter *components.ConfigurationsFilterParameters `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (o *ListConfigurationsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListConfigurationsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListConfigurationsRequest) GetFilter() *components.ConfigurationsFilterParameters {
	if o == nil {
		return nil
	}
	return o.Filter
}

type ListConfigurationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A paginated list for a collection of configurations.
	ListConfigurationsResponse *components.ListConfigurationsResponse
}

func (o *ListConfigurationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListConfigurationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListConfigurationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListConfigurationsResponse) GetListConfigurationsResponse() *components.ListConfigurationsResponse {
	if o == nil {
		return nil
	}
	return o.ListConfigurationsResponse
}
