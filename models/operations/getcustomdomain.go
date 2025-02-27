// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var GetCustomDomainServerList = []string{
	"https://global.api.konghq.com/",
}

type GetCustomDomainRequest struct {
	// ID of the custom domain to operate on.
	CustomDomainID string `pathParam:"style=simple,explode=false,name=customDomainId"`
}

func (o *GetCustomDomainRequest) GetCustomDomainID() string {
	if o == nil {
		return ""
	}
	return o.CustomDomainID
}

type GetCustomDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Response format for retrieving a custom domain for a control-plane.
	CustomDomain *components.CustomDomain
}

func (o *GetCustomDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCustomDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCustomDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCustomDomainResponse) GetCustomDomain() *components.CustomDomain {
	if o == nil {
		return nil
	}
	return o.CustomDomain
}
