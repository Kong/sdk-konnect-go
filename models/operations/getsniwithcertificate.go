// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type GetSniWithCertificateRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Certificate to lookup
	CertificateID string `pathParam:"style=simple,explode=false,name=CertificateId"`
	// ID of the SNI to lookup
	SNIID string `pathParam:"style=simple,explode=false,name=SNIId"`
}

func (o *GetSniWithCertificateRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *GetSniWithCertificateRequest) GetCertificateID() string {
	if o == nil {
		return ""
	}
	return o.CertificateID
}

func (o *GetSniWithCertificateRequest) GetSNIID() string {
	if o == nil {
		return ""
	}
	return o.SNIID
}

type GetSniWithCertificateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched SNI
	Sni *components.Sni
}

func (o *GetSniWithCertificateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSniWithCertificateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSniWithCertificateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSniWithCertificateResponse) GetSni() *components.Sni {
	if o == nil {
		return nil
	}
	return o.Sni
}
