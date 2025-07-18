// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type ReplacePortalAssetFaviconRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
	// Update an image asset for the portal.
	ReplacePortalImageAsset *components.ReplacePortalImageAsset `request:"mediaType=application/json"`
}

func (o *ReplacePortalAssetFaviconRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *ReplacePortalAssetFaviconRequest) GetReplacePortalImageAsset() *components.ReplacePortalImageAsset {
	if o == nil {
		return nil
	}
	return o.ReplacePortalImageAsset
}

type ReplacePortalAssetFaviconResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Image asset for the portal. Can be either png, jpeg or svg
	PortalAssetResponse *components.PortalAssetResponse
}

func (o *ReplacePortalAssetFaviconResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReplacePortalAssetFaviconResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReplacePortalAssetFaviconResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ReplacePortalAssetFaviconResponse) GetPortalAssetResponse() *components.PortalAssetResponse {
	if o == nil {
		return nil
	}
	return o.PortalAssetResponse
}
