// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type FetchAPIVersionRequest struct {
	// The UUID API identifier
	APIID string `pathParam:"style=simple,explode=false,name=apiId"`
	// The API version identifier
	SpecID string `pathParam:"style=simple,explode=false,name=specId"`
}

func (o *FetchAPIVersionRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *FetchAPIVersionRequest) GetSpecID() string {
	if o == nil {
		return ""
	}
	return o.SpecID
}

type FetchAPIVersionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// API version (OpenAPI or AsyncAPI)
	APIVersionResponse *components.APIVersionResponse
}

func (o *FetchAPIVersionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *FetchAPIVersionResponse) GetAPIVersionResponse() *components.APIVersionResponse {
	if o == nil {
		return nil
	}
	return o.APIVersionResponse
}
