// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

var DeleteCustomDomainServerList = []string{
	"https://global.api.konghq.com/",
}

type DeleteCustomDomainRequest struct {
	// ID of the custom domain to operate on.
	CustomDomainID string `pathParam:"style=simple,explode=false,name=customDomainId"`
}

func (o *DeleteCustomDomainRequest) GetCustomDomainID() string {
	if o == nil {
		return ""
	}
	return o.CustomDomainID
}

type DeleteCustomDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteCustomDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteCustomDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteCustomDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
