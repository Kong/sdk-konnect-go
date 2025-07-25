// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
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
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successfully created CA Certificate
	CACertificate *components.CACertificate
}

func (o *CreateCaCertificateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateCaCertificateResponse) GetCACertificate() *components.CACertificate {
	if o == nil {
		return nil
	}
	return o.CACertificate
}
