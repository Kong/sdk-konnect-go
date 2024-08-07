// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

// GatewayUnauthorizedError - Unauthorized
type GatewayUnauthorizedError struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
}

var _ error = &GatewayUnauthorizedError{}

func (e *GatewayUnauthorizedError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
