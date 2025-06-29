// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"mockserver/internal/sdk/utils"
)

type CloudGatewaysStringFieldFilterOverrideType string

const (
	CloudGatewaysStringFieldFilterOverrideTypeCloudGatewaysStringFieldEqualsFilterOverride CloudGatewaysStringFieldFilterOverrideType = "CloudGatewaysStringFieldEqualsFilterOverride"
	CloudGatewaysStringFieldFilterOverrideTypeStringFieldContainsFilter                    CloudGatewaysStringFieldFilterOverrideType = "StringFieldContainsFilter"
	CloudGatewaysStringFieldFilterOverrideTypeStringFieldNEQFilter                         CloudGatewaysStringFieldFilterOverrideType = "StringFieldNEQFilter"
	CloudGatewaysStringFieldFilterOverrideTypeStringFieldOEQFilter                         CloudGatewaysStringFieldFilterOverrideType = "StringFieldOEQFilter"
	CloudGatewaysStringFieldFilterOverrideTypeStringFieldOContainsFilter                   CloudGatewaysStringFieldFilterOverrideType = "StringFieldOContainsFilter"
)

type CloudGatewaysStringFieldFilterOverride struct {
	CloudGatewaysStringFieldEqualsFilterOverride *CloudGatewaysStringFieldEqualsFilterOverride `queryParam:"inline"`
	StringFieldContainsFilter                    *StringFieldContainsFilter                    `queryParam:"inline"`
	StringFieldNEQFilter                         *StringFieldNEQFilter                         `queryParam:"inline"`
	StringFieldOEQFilter                         *StringFieldOEQFilter                         `queryParam:"inline"`
	StringFieldOContainsFilter                   *StringFieldOContainsFilter                   `queryParam:"inline"`

	Type CloudGatewaysStringFieldFilterOverrideType
}

func CreateCloudGatewaysStringFieldFilterOverrideCloudGatewaysStringFieldEqualsFilterOverride(cloudGatewaysStringFieldEqualsFilterOverride CloudGatewaysStringFieldEqualsFilterOverride) CloudGatewaysStringFieldFilterOverride {
	typ := CloudGatewaysStringFieldFilterOverrideTypeCloudGatewaysStringFieldEqualsFilterOverride

	return CloudGatewaysStringFieldFilterOverride{
		CloudGatewaysStringFieldEqualsFilterOverride: &cloudGatewaysStringFieldEqualsFilterOverride,
		Type: typ,
	}
}

func CreateCloudGatewaysStringFieldFilterOverrideStringFieldContainsFilter(stringFieldContainsFilter StringFieldContainsFilter) CloudGatewaysStringFieldFilterOverride {
	typ := CloudGatewaysStringFieldFilterOverrideTypeStringFieldContainsFilter

	return CloudGatewaysStringFieldFilterOverride{
		StringFieldContainsFilter: &stringFieldContainsFilter,
		Type:                      typ,
	}
}

func CreateCloudGatewaysStringFieldFilterOverrideStringFieldNEQFilter(stringFieldNEQFilter StringFieldNEQFilter) CloudGatewaysStringFieldFilterOverride {
	typ := CloudGatewaysStringFieldFilterOverrideTypeStringFieldNEQFilter

	return CloudGatewaysStringFieldFilterOverride{
		StringFieldNEQFilter: &stringFieldNEQFilter,
		Type:                 typ,
	}
}

func CreateCloudGatewaysStringFieldFilterOverrideStringFieldOEQFilter(stringFieldOEQFilter StringFieldOEQFilter) CloudGatewaysStringFieldFilterOverride {
	typ := CloudGatewaysStringFieldFilterOverrideTypeStringFieldOEQFilter

	return CloudGatewaysStringFieldFilterOverride{
		StringFieldOEQFilter: &stringFieldOEQFilter,
		Type:                 typ,
	}
}

func CreateCloudGatewaysStringFieldFilterOverrideStringFieldOContainsFilter(stringFieldOContainsFilter StringFieldOContainsFilter) CloudGatewaysStringFieldFilterOverride {
	typ := CloudGatewaysStringFieldFilterOverrideTypeStringFieldOContainsFilter

	return CloudGatewaysStringFieldFilterOverride{
		StringFieldOContainsFilter: &stringFieldOContainsFilter,
		Type:                       typ,
	}
}

func (u *CloudGatewaysStringFieldFilterOverride) UnmarshalJSON(data []byte) error {

	var stringFieldContainsFilter StringFieldContainsFilter = StringFieldContainsFilter{}
	if err := utils.UnmarshalJSON(data, &stringFieldContainsFilter, "", true, true); err == nil {
		u.StringFieldContainsFilter = &stringFieldContainsFilter
		u.Type = CloudGatewaysStringFieldFilterOverrideTypeStringFieldContainsFilter
		return nil
	}

	var stringFieldNEQFilter StringFieldNEQFilter = StringFieldNEQFilter{}
	if err := utils.UnmarshalJSON(data, &stringFieldNEQFilter, "", true, true); err == nil {
		u.StringFieldNEQFilter = &stringFieldNEQFilter
		u.Type = CloudGatewaysStringFieldFilterOverrideTypeStringFieldNEQFilter
		return nil
	}

	var stringFieldOEQFilter StringFieldOEQFilter = StringFieldOEQFilter{}
	if err := utils.UnmarshalJSON(data, &stringFieldOEQFilter, "", true, true); err == nil {
		u.StringFieldOEQFilter = &stringFieldOEQFilter
		u.Type = CloudGatewaysStringFieldFilterOverrideTypeStringFieldOEQFilter
		return nil
	}

	var stringFieldOContainsFilter StringFieldOContainsFilter = StringFieldOContainsFilter{}
	if err := utils.UnmarshalJSON(data, &stringFieldOContainsFilter, "", true, true); err == nil {
		u.StringFieldOContainsFilter = &stringFieldOContainsFilter
		u.Type = CloudGatewaysStringFieldFilterOverrideTypeStringFieldOContainsFilter
		return nil
	}

	var cloudGatewaysStringFieldEqualsFilterOverride CloudGatewaysStringFieldEqualsFilterOverride = CloudGatewaysStringFieldEqualsFilterOverride{}
	if err := utils.UnmarshalJSON(data, &cloudGatewaysStringFieldEqualsFilterOverride, "", true, true); err == nil {
		u.CloudGatewaysStringFieldEqualsFilterOverride = &cloudGatewaysStringFieldEqualsFilterOverride
		u.Type = CloudGatewaysStringFieldFilterOverrideTypeCloudGatewaysStringFieldEqualsFilterOverride
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CloudGatewaysStringFieldFilterOverride", string(data))
}

func (u CloudGatewaysStringFieldFilterOverride) MarshalJSON() ([]byte, error) {
	if u.CloudGatewaysStringFieldEqualsFilterOverride != nil {
		return utils.MarshalJSON(u.CloudGatewaysStringFieldEqualsFilterOverride, "", true)
	}

	if u.StringFieldContainsFilter != nil {
		return utils.MarshalJSON(u.StringFieldContainsFilter, "", true)
	}

	if u.StringFieldNEQFilter != nil {
		return utils.MarshalJSON(u.StringFieldNEQFilter, "", true)
	}

	if u.StringFieldOEQFilter != nil {
		return utils.MarshalJSON(u.StringFieldOEQFilter, "", true)
	}

	if u.StringFieldOContainsFilter != nil {
		return utils.MarshalJSON(u.StringFieldOContainsFilter, "", true)
	}

	return nil, errors.New("could not marshal union type CloudGatewaysStringFieldFilterOverride: all fields are null")
}
