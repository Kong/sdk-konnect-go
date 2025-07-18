// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"mockserver/internal/sdk/models/components"
)

// KonnectCPLegacyNotFoundError - standard error
type KonnectCPLegacyNotFoundError struct {
	Message  any                     `json:"message,omitempty"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &KonnectCPLegacyNotFoundError{}

func (e *KonnectCPLegacyNotFoundError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
