// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

// RateLimited - The error object
type RateLimited struct {
	// The HTTP response code
	Status *int64 `json:"status,omitempty"`
	// The Error response
	Title *string `json:"title,omitempty"`
	// The Konnect traceback ID.
	Instance *string `json:"instance,omitempty"`
	// Detailed explanation of the error response.
	Detail *string `json:"detail,omitempty"`
}

var _ error = &RateLimited{}

func (e *RateLimited) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}