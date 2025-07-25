// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type FetchAPIVersionRequest struct {
	// The UUID API identifier
	APIID string `pathParam:"style=simple,explode=false,name=apiId"`
	// The API version identifier
	VersionID string `pathParam:"style=simple,explode=false,name=versionId"`
}

func (o *FetchAPIVersionRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *FetchAPIVersionRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type FetchAPIVersionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// API version (OpenAPI or AsyncAPI)
	APIVersionResponse *components.APIVersionResponse
}

func (o *FetchAPIVersionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *FetchAPIVersionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *FetchAPIVersionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *FetchAPIVersionResponse) GetAPIVersionResponse() *components.APIVersionResponse {
	if o == nil {
		return nil
	}
	return o.APIVersionResponse
}
