// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type DeleteAppAuthStrategyRequest struct {
	// Application auth strategy identifier
	AuthStrategyID string `pathParam:"style=simple,explode=false,name=authStrategyId"`
}

func (o *DeleteAppAuthStrategyRequest) GetAuthStrategyID() string {
	if o == nil {
		return ""
	}
	return o.AuthStrategyID
}

type DeleteAppAuthStrategyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteAppAuthStrategyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
