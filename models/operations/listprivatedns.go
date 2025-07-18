// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var ListPrivateDNSServerList = []string{
	"https://global.api.konghq.com/",
}

type ListPrivateDNSRequest struct {
	// The network to operate on.
	NetworkID string `pathParam:"style=simple,explode=false,name=networkId"`
	// Filters supported for Private DNS.
	Filter *components.PrivateDNSFilterParameters `queryParam:"style=deepObject,explode=true,name=filter"`
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
}

func (o *ListPrivateDNSRequest) GetNetworkID() string {
	if o == nil {
		return ""
	}
	return o.NetworkID
}

func (o *ListPrivateDNSRequest) GetFilter() *components.PrivateDNSFilterParameters {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *ListPrivateDNSRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListPrivateDNSRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

type ListPrivateDNSResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A paginated list for a collection of Private DNS for a network.
	ListPrivateDNSResponse *components.ListPrivateDNSResponse
}

func (o *ListPrivateDNSResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPrivateDNSResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPrivateDNSResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListPrivateDNSResponse) GetListPrivateDNSResponse() *components.ListPrivateDNSResponse {
	if o == nil {
		return nil
	}
	return o.ListPrivateDNSResponse
}
