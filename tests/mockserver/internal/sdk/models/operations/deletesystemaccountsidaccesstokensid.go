// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type DeleteSystemAccountsIDAccessTokensIDRequest struct {
	// ID of the system account.
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
	// ID of the system account access token.
	TokenID string `pathParam:"style=simple,explode=false,name=tokenId"`
}

func (o *DeleteSystemAccountsIDAccessTokensIDRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *DeleteSystemAccountsIDAccessTokensIDRequest) GetTokenID() string {
	if o == nil {
		return ""
	}
	return o.TokenID
}

type DeleteSystemAccountsIDAccessTokensIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteSystemAccountsIDAccessTokensIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
