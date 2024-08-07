// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/utils"
)

type StringFieldFilterType string

const (
	StringFieldFilterTypeStr                       StringFieldFilterType = "str"
	StringFieldFilterTypeStringFieldContainsFilter StringFieldFilterType = "StringFieldContainsFilter"
)

// StringFieldFilter - Filter a string value field either by exact match or partial contains.
type StringFieldFilter struct {
	Str                       *string
	StringFieldContainsFilter *StringFieldContainsFilter

	Type StringFieldFilterType
}

func CreateStringFieldFilterStr(str string) StringFieldFilter {
	typ := StringFieldFilterTypeStr

	return StringFieldFilter{
		Str:  &str,
		Type: typ,
	}
}

func CreateStringFieldFilterStringFieldContainsFilter(stringFieldContainsFilter StringFieldContainsFilter) StringFieldFilter {
	typ := StringFieldFilterTypeStringFieldContainsFilter

	return StringFieldFilter{
		StringFieldContainsFilter: &stringFieldContainsFilter,
		Type:                      typ,
	}
}

func (u *StringFieldFilter) UnmarshalJSON(data []byte) error {

	var stringFieldContainsFilter StringFieldContainsFilter = StringFieldContainsFilter{}
	if err := utils.UnmarshalJSON(data, &stringFieldContainsFilter, "", true, true); err == nil {
		u.StringFieldContainsFilter = &stringFieldContainsFilter
		u.Type = StringFieldFilterTypeStringFieldContainsFilter
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = StringFieldFilterTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for StringFieldFilter", string(data))
}

func (u StringFieldFilter) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.StringFieldContainsFilter != nil {
		return utils.MarshalJSON(u.StringFieldContainsFilter, "", true)
	}

	return nil, errors.New("could not marshal union type StringFieldFilter: all fields are null")
}
