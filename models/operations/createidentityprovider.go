// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var CreateIdentityProviderServerList = []string{
	"https://global.api.konghq.com/",
}

type CreateIdentityProviderResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// An identity provider configuration. This response represents the configuration of a specific identity provider, which can be either OIDC or SAML.
	//
	IdentityProvider *components.IdentityProvider
}

func (o *CreateIdentityProviderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateIdentityProviderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateIdentityProviderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateIdentityProviderResponse) GetIdentityProvider() *components.IdentityProvider {
	if o == nil {
		return nil
	}
	return o.IdentityProvider
}
