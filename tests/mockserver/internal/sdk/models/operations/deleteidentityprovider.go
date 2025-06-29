// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type DeleteIdentityProviderRequest struct {
	// ID of the identity provider.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteIdentityProviderRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteIdentityProviderResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteIdentityProviderResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
