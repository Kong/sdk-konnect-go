// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

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
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A paginated list for a collection of configurations.
	ListConfigurationsResponse *components.ListConfigurationsResponse
}

func (o *ListConfigurationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListConfigurationsResponse) GetListConfigurationsResponse() *components.ListConfigurationsResponse {
	if o == nil {
		return nil
	}
	return o.ListConfigurationsResponse
}
