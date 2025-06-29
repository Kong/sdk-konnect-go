// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
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
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A response including a single system account.
	SystemAccount *components.SystemAccount
}

func (o *PatchSystemAccountsIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PatchSystemAccountsIDResponse) GetSystemAccount() *components.SystemAccount {
	if o == nil {
		return nil
	}
	return o.SystemAccount
}
