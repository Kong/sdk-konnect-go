// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type CreateCaCertificateRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Description of the new CA Certificate for creation
	CACertificate components.CACertificate `request:"mediaType=application/json"`
}

func (o *CreateCaCertificateRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateCaCertificateRequest) GetCACertificate() components.CACertificate {
	if o == nil {
		return components.CACertificate{}
	}
	return o.CACertificate
}

type CreateCaCertificateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created CA Certificate
	CACertificate *components.CACertificate
}

func (o *CreateCaCertificateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCaCertificateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCaCertificateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateCaCertificateResponse) GetCACertificate() *components.CACertificate {
	if o == nil {
		return nil
	}
	return o.CACertificate
}
