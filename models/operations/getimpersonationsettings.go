// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var GetImpersonationSettingsServerList = []string{
	"https://global.api.konghq.com/",
}

type GetImpersonationSettingsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Response for Get Impersonation Settings endpoint
	GetImpersonationSettingsResponse *components.GetImpersonationSettingsResponse
}

func (o *GetImpersonationSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetImpersonationSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetImpersonationSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetImpersonationSettingsResponse) GetGetImpersonationSettingsResponse() *components.GetImpersonationSettingsResponse {
	if o == nil {
		return nil
	}
	return o.GetImpersonationSettingsResponse
}
