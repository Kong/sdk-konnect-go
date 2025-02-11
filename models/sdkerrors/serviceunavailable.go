// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// Status - The HTTP status code.
type Status int

const (
	StatusFiveHundredAndThree Status = 503
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 503:
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

// ServiceUnavailable - Error response for temporary service unavailability.
type ServiceUnavailable struct {
	// The HTTP status code.
	Status Status `json:"status"`
	// The error response code.
	Title string `json:"title"`
	// The Konnect traceback code
	Instance string `json:"instance"`
	// Details about the error.
	Detail *string `json:"detail,omitempty"`
}

var _ error = &ServiceUnavailable{}

func (e *ServiceUnavailable) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
