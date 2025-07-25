// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type PostPortalCustomEmailTestSendRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
	// Name of the email template.
	TemplateName components.EmailTemplateName `pathParam:"style=simple,explode=false,name=templateName"`
	// Send a test email using a custom email template.
	PostSendTestEmailPayload components.PostSendTestEmailPayload `request:"mediaType=application/json"`
}

func (o *PostPortalCustomEmailTestSendRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *PostPortalCustomEmailTestSendRequest) GetTemplateName() components.EmailTemplateName {
	if o == nil {
		return components.EmailTemplateName("")
	}
	return o.TemplateName
}

func (o *PostPortalCustomEmailTestSendRequest) GetPostSendTestEmailPayload() components.PostSendTestEmailPayload {
	if o == nil {
		return components.PostSendTestEmailPayload{}
	}
	return o.PostSendTestEmailPayload
}

type PostPortalCustomEmailTestSendResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostPortalCustomEmailTestSendResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostPortalCustomEmailTestSendResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostPortalCustomEmailTestSendResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
