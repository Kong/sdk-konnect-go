// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type ListResourceQuotasRequest struct {
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
}

func (o *ListResourceQuotasRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListResourceQuotasRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

type ListResourceQuotasResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A paginated list for a collection of resource quotas.
	ListResourceQuotasResponse *components.ListResourceQuotasResponse
}

func (o *ListResourceQuotasResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListResourceQuotasResponse) GetListResourceQuotasResponse() *components.ListResourceQuotasResponse {
	if o == nil {
		return nil
	}
	return o.ListResourceQuotasResponse
}
