// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"io"
	"mockserver/internal/sdk/models/components"
)

type GetPortalAssetFaviconRawRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
}

func (o *GetPortalAssetFaviconRawRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

type GetPortalAssetFaviconRawResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Logo of the portal. Can be either png, jpeg or svg
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	TwoHundredImagePngPortalImageAssetBlob io.ReadCloser
	// Logo of the portal. Can be either png, jpeg or svg
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	TwoHundredImageJpegPortalImageAssetBlob io.ReadCloser
	// Logo of the portal. Can be either png, jpeg or svg
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	TwoHundredImageSvgPlusXMLPortalImageAssetBlob io.ReadCloser
}

func (o *GetPortalAssetFaviconRawResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPortalAssetFaviconRawResponse) GetTwoHundredImagePngPortalImageAssetBlob() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.TwoHundredImagePngPortalImageAssetBlob
}

func (o *GetPortalAssetFaviconRawResponse) GetTwoHundredImageJpegPortalImageAssetBlob() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.TwoHundredImageJpegPortalImageAssetBlob
}

func (o *GetPortalAssetFaviconRawResponse) GetTwoHundredImageSvgPlusXMLPortalImageAssetBlob() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.TwoHundredImageSvgPlusXMLPortalImageAssetBlob
}
