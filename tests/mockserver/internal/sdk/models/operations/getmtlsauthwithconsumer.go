// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetMtlsAuthWithConsumerRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Consumer ID for nested entities
	ConsumerIDForNestedEntities string `pathParam:"style=simple,explode=false,name=ConsumerIdForNestedEntities"`
	// ID of the MTLS-auth credential to lookup
	MTLSAuthID string `pathParam:"style=simple,explode=false,name=MTLSAuthId"`
}

func (o *GetMtlsAuthWithConsumerRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *GetMtlsAuthWithConsumerRequest) GetConsumerIDForNestedEntities() string {
	if o == nil {
		return ""
	}
	return o.ConsumerIDForNestedEntities
}

func (o *GetMtlsAuthWithConsumerRequest) GetMTLSAuthID() string {
	if o == nil {
		return ""
	}
	return o.MTLSAuthID
}

type GetMtlsAuthWithConsumerResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successfully fetched MTLS-auth credential
	MTLSAuth *components.MTLSAuth
}

func (o *GetMtlsAuthWithConsumerResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMtlsAuthWithConsumerResponse) GetMTLSAuth() *components.MTLSAuth {
	if o == nil {
		return nil
	}
	return o.MTLSAuth
}
