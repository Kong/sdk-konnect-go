// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

// PreconditionFailed - The error response object.
type PreconditionFailed struct {
	// The HTTP status code.
	Status *int64 `json:"status,omitempty"`
	// The error response code.
	Title *string `json:"title,omitempty"`
	// The Konnect traceback code.
	Instance *string `json:"instance,omitempty"`
	// Details about the error response.
	Detail *string `json:"detail,omitempty"`
}

var _ error = &PreconditionFailed{}

func (e *PreconditionFailed) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
