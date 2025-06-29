// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type PostSystemAccountsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A response including a single system account.
	SystemAccount *components.SystemAccount
}

func (o *PostSystemAccountsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostSystemAccountsResponse) GetSystemAccount() *components.SystemAccount {
	if o == nil {
		return nil
	}
	return o.SystemAccount
}
