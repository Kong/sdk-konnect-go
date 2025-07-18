// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type ListPortalsRequest struct {
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Sorts a collection of portals. Supported sort attributes are:
	//   - name
	//   - description
	//   - authentication_enabled
	//   - rbac_enabled
	//   - auto_approve_applications
	//   - auto_approve_developers
	//   - default_domain
	//   - canonical_domain
	//   - created_at
	//   - updated_at
	//
	Sort *string `queryParam:"style=form,explode=true,name=sort"`
	// Filter portals returned in the response.
	Filter *components.PortalFilterParameters `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (o *ListPortalsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListPortalsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListPortalsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListPortalsRequest) GetFilter() *components.PortalFilterParameters {
	if o == nil {
		return nil
	}
	return o.Filter
}

type ListPortalsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A paginated list of portals in the current region of an organization.
	ListPortalsResponse *components.ListPortalsResponse
}

func (o *ListPortalsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPortalsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPortalsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListPortalsResponse) GetListPortalsResponse() *components.ListPortalsResponse {
	if o == nil {
		return nil
	}
	return o.ListPortalsResponse
}
