// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type UpdateAPIVersionRequest struct {
	// The UUID API identifier
	APIID string `pathParam:"style=simple,explode=false,name=apiId"`
	// The API version identifier
	SpecID     string                `pathParam:"style=simple,explode=false,name=specId"`
	APIVersion components.APIVersion `request:"mediaType=application/json"`
}

func (o *UpdateAPIVersionRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *UpdateAPIVersionRequest) GetSpecID() string {
	if o == nil {
		return ""
	}
	return o.SpecID
}

func (o *UpdateAPIVersionRequest) GetAPIVersion() components.APIVersion {
	if o == nil {
		return components.APIVersion{}
	}
	return o.APIVersion
}

type UpdateAPIVersionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// API version (OpenAPI or AsyncAPI)
	APIVersionResponse *components.APIVersionResponse
}

func (o *UpdateAPIVersionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateAPIVersionResponse) GetAPIVersionResponse() *components.APIVersionResponse {
	if o == nil {
		return nil
	}
	return o.APIVersionResponse
}
