// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var GetAuthenticationSettingsServerList = []string{
	"https://global.api.konghq.com/",
}

type GetAuthenticationSettingsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Response for authentication settings endpoint
	AuthenticationSettings *components.AuthenticationSettings
}

func (o *GetAuthenticationSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAuthenticationSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAuthenticationSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAuthenticationSettingsResponse) GetAuthenticationSettings() *components.AuthenticationSettings {
	if o == nil {
		return nil
	}
	return o.AuthenticationSettings
}
