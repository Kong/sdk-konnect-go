// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// InvalidRules - invalid parameters rules
type InvalidRules string

const (
	InvalidRulesRequired                               InvalidRules = "required"
	InvalidRulesIsArray                                InvalidRules = "is_array"
	InvalidRulesIsBase64                               InvalidRules = "is_base64"
	InvalidRulesIsBoolean                              InvalidRules = "is_boolean"
	InvalidRulesIsDateTime                             InvalidRules = "is_date_time"
	InvalidRulesIsInteger                              InvalidRules = "is_integer"
	InvalidRulesIsNull                                 InvalidRules = "is_null"
	InvalidRulesIsNumber                               InvalidRules = "is_number"
	InvalidRulesIsObject                               InvalidRules = "is_object"
	InvalidRulesIsString                               InvalidRules = "is_string"
	InvalidRulesIsUUID                                 InvalidRules = "is_uuid"
	InvalidRulesIsFqdn                                 InvalidRules = "is_fqdn"
	InvalidRulesIsArn                                  InvalidRules = "is_arn"
	InvalidRulesUnknownProperty                        InvalidRules = "unknown_property"
	InvalidRulesMissingReference                       InvalidRules = "missing_reference"
	InvalidRulesIsLabel                                InvalidRules = "is_label"
	InvalidRulesMatchesRegex                           InvalidRules = "matches_regex"
	InvalidRulesInvalid                                InvalidRules = "invalid"
	InvalidRulesIsSupportedNetworkAvailabilityZoneList InvalidRules = "is_supported_network_availability_zone_list"
	InvalidRulesIsSupportedNetworkCidrBlock            InvalidRules = "is_supported_network_cidr_block"
	InvalidRulesIsSupportedProviderRegion              InvalidRules = "is_supported_provider_region"
)

func (e InvalidRules) ToPointer() *InvalidRules {
	return &e
}
func (e *InvalidRules) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "required":
		fallthrough
	case "is_array":
		fallthrough
	case "is_base64":
		fallthrough
	case "is_boolean":
		fallthrough
	case "is_date_time":
		fallthrough
	case "is_integer":
		fallthrough
	case "is_null":
		fallthrough
	case "is_number":
		fallthrough
	case "is_object":
		fallthrough
	case "is_string":
		fallthrough
	case "is_uuid":
		fallthrough
	case "is_fqdn":
		fallthrough
	case "is_arn":
		fallthrough
	case "unknown_property":
		fallthrough
	case "missing_reference":
		fallthrough
	case "is_label":
		fallthrough
	case "matches_regex":
		fallthrough
	case "invalid":
		fallthrough
	case "is_supported_network_availability_zone_list":
		fallthrough
	case "is_supported_network_cidr_block":
		fallthrough
	case "is_supported_provider_region":
		*e = InvalidRules(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InvalidRules: %v", v)
	}
}
