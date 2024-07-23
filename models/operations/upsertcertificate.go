// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type UpsertCertificateRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Certificate to lookup
	CertificateID string `pathParam:"style=simple,explode=false,name=CertificateId"`
	// Description of the Certificate
	Certificate components.CertificateInput `request:"mediaType=application/json"`
}

func (o *UpsertCertificateRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpsertCertificateRequest) GetCertificateID() string {
	if o == nil {
		return ""
	}
	return o.CertificateID
}

func (o *UpsertCertificateRequest) GetCertificate() components.CertificateInput {
	if o == nil {
		return components.CertificateInput{}
	}
	return o.Certificate
}

type UpsertCertificateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted Certificate
	Certificate *components.Certificate
}

func (o *UpsertCertificateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertCertificateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertCertificateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertCertificateResponse) GetCertificate() *components.Certificate {
	if o == nil {
		return nil
	}
	return o.Certificate
}
