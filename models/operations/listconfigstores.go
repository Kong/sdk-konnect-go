// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type ListConfigStoresRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Request the next page of data, starting with the item after this parameter.
	PageAfter *string `queryParam:"style=form,explode=true,name=page[after]"`
}

func (o *ListConfigStoresRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *ListConfigStoresRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListConfigStoresRequest) GetPageAfter() *string {
	if o == nil {
		return nil
	}
	return o.PageAfter
}

type ListConfigStoresResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List of Config Stores
	ListConfigStoresResponse *components.ListConfigStoresResponse
}

func (o *ListConfigStoresResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListConfigStoresResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListConfigStoresResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListConfigStoresResponse) GetListConfigStoresResponse() *components.ListConfigStoresResponse {
	if o == nil {
		return nil
	}
	return o.ListConfigStoresResponse
}
