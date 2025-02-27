// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type ListUserConfigurationsRequest struct {
	// Filters a collection of configurations.
	Filter *components.ConfigurationFilterParameters `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (o *ListUserConfigurationsRequest) GetFilter() *components.ConfigurationFilterParameters {
	if o == nil {
		return nil
	}
	return o.Filter
}

type ListUserConfigurationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A paginated list response for all notification events and user subscriptions.
	UserConfigurationListResponse *components.UserConfigurationListResponse
}

func (o *ListUserConfigurationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUserConfigurationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUserConfigurationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListUserConfigurationsResponse) GetUserConfigurationListResponse() *components.UserConfigurationListResponse {
	if o == nil {
		return nil
	}
	return o.UserConfigurationListResponse
}
