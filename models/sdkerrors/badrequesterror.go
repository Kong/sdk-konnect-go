// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"github.com/Kong/sdk-konnect-go/models/components"
)

// BadRequestError - standard error
type BadRequestError struct {
	// The HTTP status code of the error. Useful when passing the response
	// body to child properties in a frontend UI. Must be returned as an integer.
	//
	Status int64 `json:"status"`
	// A short, human-readable summary of the problem. It should not
	// change between occurences of a problem, except for localization.
	// Should be provided as "Sentence case" for direct use in the UI.
	//
	Title string `json:"title"`
	// The error type.
	Type *string `json:"type,omitempty"`
	// Used to return the correlation ID back to the user, in the format
	// kong:trace:<correlation_id>. This helps us find the relevant logs
	// when a customer reports an issue.
	//
	Instance string `json:"instance"`
	// A human readable explanation specific to this occurence of the problem.
	// This field may contain request/entity data to help the user understand
	// what went wrong. Enclose variable values in square brackets. Should be
	// provided as "Sentence case" for direct use in the UI.
	//
	Detail string `json:"detail"`
	// invalid parameters
	InvalidParameters []components.InvalidParameters `json:"invalid_parameters"`
}

var _ error = &BadRequestError{}

func (e *BadRequestError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
