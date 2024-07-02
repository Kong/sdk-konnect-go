// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type PatchSystemAccountsIDRequest struct {
	// ID of the system account.
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
	// The request schema for the update system account request.
	UpdateSystemAccount *components.UpdateSystemAccount `request:"mediaType=application/json"`
}

func (o *PatchSystemAccountsIDRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *PatchSystemAccountsIDRequest) GetUpdateSystemAccount() *components.UpdateSystemAccount {
	if o == nil {
		return nil
	}
	return o.UpdateSystemAccount
}

type PatchSystemAccountsIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response including a single system account.
	SystemAccount *components.SystemAccount
}

func (o *PatchSystemAccountsIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchSystemAccountsIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchSystemAccountsIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchSystemAccountsIDResponse) GetSystemAccount() *components.SystemAccount {
	if o == nil {
		return nil
	}
	return o.SystemAccount
}
