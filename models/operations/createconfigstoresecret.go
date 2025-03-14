// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type CreateConfigStoreSecretRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Config Store identifier
	ConfigStoreID           string                             `pathParam:"style=simple,explode=false,name=configStoreId"`
	CreateConfigStoreSecret components.CreateConfigStoreSecret `request:"mediaType=application/json"`
}

func (o *CreateConfigStoreSecretRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateConfigStoreSecretRequest) GetConfigStoreID() string {
	if o == nil {
		return ""
	}
	return o.ConfigStoreID
}

func (o *CreateConfigStoreSecretRequest) GetCreateConfigStoreSecret() components.CreateConfigStoreSecret {
	if o == nil {
		return components.CreateConfigStoreSecret{}
	}
	return o.CreateConfigStoreSecret
}

type CreateConfigStoreSecretResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Config Store Secret
	ConfigStoreSecret *components.ConfigStoreSecret
}

func (o *CreateConfigStoreSecretResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateConfigStoreSecretResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateConfigStoreSecretResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateConfigStoreSecretResponse) GetConfigStoreSecret() *components.ConfigStoreSecret {
	if o == nil {
		return nil
	}
	return o.ConfigStoreSecret
}
