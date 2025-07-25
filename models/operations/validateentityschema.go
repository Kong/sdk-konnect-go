// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type ValidateEntitySchemaRequest struct {
	// The name of the entity
	EntityName string `pathParam:"style=simple,explode=false,name=entityName"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Request body of a Koko entity to validate against its schema
	RequestBody map[string]any `request:"mediaType=application/json"`
}

func (o *ValidateEntitySchemaRequest) GetEntityName() string {
	if o == nil {
		return ""
	}
	return o.EntityName
}

func (o *ValidateEntitySchemaRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *ValidateEntitySchemaRequest) GetRequestBody() map[string]any {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type ValidateEntitySchemaResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Validation result of the request against a schema
	ValidateEntityResponse *components.ValidateEntityResponse
}

func (o *ValidateEntitySchemaResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ValidateEntitySchemaResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ValidateEntitySchemaResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ValidateEntitySchemaResponse) GetValidateEntityResponse() *components.ValidateEntityResponse {
	if o == nil {
		return nil
	}
	return o.ValidateEntityResponse
}
