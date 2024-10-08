// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type ListDataplaneNodesRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Request the next page of data, starting with the item after this parameter.
	PageAfter *string `queryParam:"style=form,explode=true,name=page[after]"`
}

func (o *ListDataplaneNodesRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *ListDataplaneNodesRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListDataplaneNodesRequest) GetPageAfter() *string {
	if o == nil {
		return nil
	}
	return o.PageAfter
}

type ListDataplaneNodesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Example response
	ListNodes *components.ListNodes
}

func (o *ListDataplaneNodesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListDataplaneNodesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListDataplaneNodesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListDataplaneNodesResponse) GetListNodes() *components.ListNodes {
	if o == nil {
		return nil
	}
	return o.ListNodes
}
