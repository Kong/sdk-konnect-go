// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

var DeleteNetworkServerList = []string{
	"https://global.api.konghq.com/",
}

type DeleteNetworkRequest struct {
	// The network to operate on.
	NetworkID string `pathParam:"style=simple,explode=false,name=networkId"`
}

func (o *DeleteNetworkRequest) GetNetworkID() string {
	if o == nil {
		return ""
	}
	return o.NetworkID
}

type DeleteNetworkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteNetworkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteNetworkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteNetworkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
