// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type ListControlPlanesRequest struct {
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Filter by direct equality comparison of the name property with a supplied value.
	FilterNameEq *string `queryParam:"style=form,explode=true,name=filter[name][eq]"`
	// Filter by direct equality comparison (short-hand) of the name property with a supplied value.
	FilterName *string `queryParam:"style=form,explode=true,name=filter[name]"`
	// Filter by contains comparison of the name property with a supplied substring.
	FilterNameContains *string `queryParam:"style=form,explode=true,name=filter[name][contains]"`
	// Filter by non-equality comparison of the name property with a supplied value.
	FilterNameNeq *string `queryParam:"style=form,explode=true,name=filter[name][neq]"`
	// Filter by direct equality comparison of the id property with a supplied value.
	FilterIDEq *string `queryParam:"style=form,explode=true,name=filter[id][eq]"`
	// Filter by direct equality comparison (short-hand) of the id property with a supplied value.
	FilterID *string `queryParam:"style=form,explode=true,name=filter[id]"`
	// Filter by direct equality comparison of id property with multiple supplied values.
	FilterIDOeq *string `queryParam:"style=form,explode=true,name=filter[id][oeq]"`
	// Filter by direct equality comparison of the cluster_type property with a supplied value.
	FilterClusterTypeEq *string `queryParam:"style=form,explode=true,name=filter[cluster_type][eq]"`
	// Filter by direct equality comparison (short-hand) of the cluster_type property with a supplied value.
	FilterClusterType *string `queryParam:"style=form,explode=true,name=filter[cluster_type]"`
	// Filter by non-equality comparison of the cluster_type property with a supplied value.
	FilterClusterTypeNeq *string `queryParam:"style=form,explode=true,name=filter[cluster_type][neq]"`
	// Filter control planes in the response by associated labels.
	Labels *string `queryParam:"style=form,explode=true,name=labels"`
}

func (o *ListControlPlanesRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListControlPlanesRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListControlPlanesRequest) GetFilterNameEq() *string {
	if o == nil {
		return nil
	}
	return o.FilterNameEq
}

func (o *ListControlPlanesRequest) GetFilterName() *string {
	if o == nil {
		return nil
	}
	return o.FilterName
}

func (o *ListControlPlanesRequest) GetFilterNameContains() *string {
	if o == nil {
		return nil
	}
	return o.FilterNameContains
}

func (o *ListControlPlanesRequest) GetFilterNameNeq() *string {
	if o == nil {
		return nil
	}
	return o.FilterNameNeq
}

func (o *ListControlPlanesRequest) GetFilterIDEq() *string {
	if o == nil {
		return nil
	}
	return o.FilterIDEq
}

func (o *ListControlPlanesRequest) GetFilterID() *string {
	if o == nil {
		return nil
	}
	return o.FilterID
}

func (o *ListControlPlanesRequest) GetFilterIDOeq() *string {
	if o == nil {
		return nil
	}
	return o.FilterIDOeq
}

func (o *ListControlPlanesRequest) GetFilterClusterTypeEq() *string {
	if o == nil {
		return nil
	}
	return o.FilterClusterTypeEq
}

func (o *ListControlPlanesRequest) GetFilterClusterType() *string {
	if o == nil {
		return nil
	}
	return o.FilterClusterType
}

func (o *ListControlPlanesRequest) GetFilterClusterTypeNeq() *string {
	if o == nil {
		return nil
	}
	return o.FilterClusterTypeNeq
}

func (o *ListControlPlanesRequest) GetLabels() *string {
	if o == nil {
		return nil
	}
	return o.Labels
}

type ListControlPlanesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A paginated list response for a collection of control planes.
	ListControlPlanesResponse *components.ListControlPlanesResponse
}

func (o *ListControlPlanesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListControlPlanesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListControlPlanesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListControlPlanesResponse) GetListControlPlanesResponse() *components.ListControlPlanesResponse {
	if o == nil {
		return nil
	}
	return o.ListControlPlanesResponse
}