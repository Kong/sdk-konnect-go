// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var UpdateNetworkServerList = []string{
	"https://global.api.konghq.com/",
}

type UpdateNetworkRequest struct {
	// The network to operate on.
	NetworkID           string                         `pathParam:"style=simple,explode=false,name=networkId"`
	PatchNetworkRequest components.PatchNetworkRequest `request:"mediaType=application/json"`
}

func (o *UpdateNetworkRequest) GetNetworkID() string {
	if o == nil {
		return ""
	}
	return o.NetworkID
}

func (o *UpdateNetworkRequest) GetPatchNetworkRequest() components.PatchNetworkRequest {
	if o == nil {
		return components.PatchNetworkRequest{}
	}
	return o.PatchNetworkRequest
}

type UpdateNetworkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Response format for patching a network.
	Network *components.Network
}

func (o *UpdateNetworkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateNetworkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateNetworkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateNetworkResponse) GetNetwork() *components.Network {
	if o == nil {
		return nil
	}
	return o.Network
}
