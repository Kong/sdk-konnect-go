// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// HighVelocityCloudGatewaysProviderName - Name of cloud provider.
type HighVelocityCloudGatewaysProviderName string

const (
	HighVelocityCloudGatewaysProviderNameAws HighVelocityCloudGatewaysProviderName = "aws"
)

func (e HighVelocityCloudGatewaysProviderName) ToPointer() *HighVelocityCloudGatewaysProviderName {
	return &e
}
func (e *HighVelocityCloudGatewaysProviderName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "aws":
		*e = HighVelocityCloudGatewaysProviderName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HighVelocityCloudGatewaysProviderName: %v", v)
	}
}
