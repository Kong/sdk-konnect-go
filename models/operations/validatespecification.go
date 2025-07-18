// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type ValidateSpecificationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// API specification (OpenAPI or AsyncAPI) validation successful
	ValidateAPISpecSuccessResponse *components.ValidateAPISpecSuccessResponse
}

func (o *ValidateSpecificationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ValidateSpecificationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ValidateSpecificationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ValidateSpecificationResponse) GetValidateAPISpecSuccessResponse() *components.ValidateAPISpecSuccessResponse {
	if o == nil {
		return nil
	}
	return o.ValidateAPISpecSuccessResponse
}
