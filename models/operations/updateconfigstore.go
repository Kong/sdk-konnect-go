// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type UpdateConfigStoreRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Config Store identifier
	ConfigStoreID     string                       `pathParam:"style=simple,explode=false,name=configStoreId"`
	UpdateConfigStore components.UpdateConfigStore `request:"mediaType=application/json"`
}

func (o *UpdateConfigStoreRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateConfigStoreRequest) GetConfigStoreID() string {
	if o == nil {
		return ""
	}
	return o.ConfigStoreID
}

func (o *UpdateConfigStoreRequest) GetUpdateConfigStore() components.UpdateConfigStore {
	if o == nil {
		return components.UpdateConfigStore{}
	}
	return o.UpdateConfigStore
}

type UpdateConfigStoreResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Config Store
	ConfigStore *components.ConfigStore
}

func (o *UpdateConfigStoreResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateConfigStoreResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateConfigStoreResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateConfigStoreResponse) GetConfigStore() *components.ConfigStore {
	if o == nil {
		return nil
	}
	return o.ConfigStore
}
