// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ProviderName - Name of cloud provider.
type ProviderName string

const (
	ProviderNameAws   ProviderName = "aws"
	ProviderNameAzure ProviderName = "azure"
)

func (e ProviderName) ToPointer() *ProviderName {
	return &e
}
func (e *ProviderName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "aws":
		fallthrough
	case "azure":
		*e = ProviderName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProviderName: %v", v)
	}
}
