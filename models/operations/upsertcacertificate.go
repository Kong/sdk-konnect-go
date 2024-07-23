// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type UpsertCaCertificateRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the CA Certificate to lookup
	CACertificateID string `pathParam:"style=simple,explode=false,name=CACertificateId"`
	// Description of the CA Certificate
	CACertificate components.CACertificateInput `request:"mediaType=application/json"`
}

func (o *UpsertCaCertificateRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpsertCaCertificateRequest) GetCACertificateID() string {
	if o == nil {
		return ""
	}
	return o.CACertificateID
}

func (o *UpsertCaCertificateRequest) GetCACertificate() components.CACertificateInput {
	if o == nil {
		return components.CACertificateInput{}
	}
	return o.CACertificate
}

type UpsertCaCertificateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted CA Certificate
	CACertificate *components.CACertificate
}

func (o *UpsertCaCertificateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertCaCertificateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertCaCertificateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertCaCertificateResponse) GetCACertificate() *components.CACertificate {
	if o == nil {
		return nil
	}
	return o.CACertificate
}
