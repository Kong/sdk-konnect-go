// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type MovePortalPagesRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
	// ID of the page.
	PageID string `pathParam:"style=simple,explode=false,name=pageId"`
	// move page
	MovePageRequestPayload *components.MovePageRequestPayload `request:"mediaType=application/json"`
}

func (o *MovePortalPagesRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *MovePortalPagesRequest) GetPageID() string {
	if o == nil {
		return ""
	}
	return o.PageID
}

func (o *MovePortalPagesRequest) GetMovePageRequestPayload() *components.MovePageRequestPayload {
	if o == nil {
		return nil
	}
	return o.MovePageRequestPayload
}

type MovePortalPagesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *MovePortalPagesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MovePortalPagesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MovePortalPagesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
