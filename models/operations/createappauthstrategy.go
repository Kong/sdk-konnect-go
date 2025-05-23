// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type CreateAppAuthStrategyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response containing the newly created application auth strategy object.
	CreateAppAuthStrategyResponse *components.CreateAppAuthStrategyResponse
}

func (o *CreateAppAuthStrategyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAppAuthStrategyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAppAuthStrategyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateAppAuthStrategyResponse) GetCreateAppAuthStrategyResponse() *components.CreateAppAuthStrategyResponse {
	if o == nil {
		return nil
	}
	return o.CreateAppAuthStrategyResponse
}

func (o *CreateAppAuthStrategyResponse) GetCreateAppAuthStrategyResponseKeyAuth() *components.AppAuthStrategyKeyAuthResponse {
	if v := o.GetCreateAppAuthStrategyResponse(); v != nil {
		return v.AppAuthStrategyKeyAuthResponse
	}
	return nil
}

func (o *CreateAppAuthStrategyResponse) GetCreateAppAuthStrategyResponseOpenidConnect() *components.AppAuthStrategyOpenIDConnectResponse {
	if v := o.GetCreateAppAuthStrategyResponse(); v != nil {
		return v.AppAuthStrategyOpenIDConnectResponse
	}
	return nil
}
