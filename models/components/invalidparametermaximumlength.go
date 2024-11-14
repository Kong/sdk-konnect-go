// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// InvalidParameterMaximumLengthRule - invalid parameters rules
type InvalidParameterMaximumLengthRule string

const (
	InvalidParameterMaximumLengthRuleMaxLength InvalidParameterMaximumLengthRule = "max_length"
	InvalidParameterMaximumLengthRuleMaxItems  InvalidParameterMaximumLengthRule = "max_items"
	InvalidParameterMaximumLengthRuleMax       InvalidParameterMaximumLengthRule = "max"
)

func (e InvalidParameterMaximumLengthRule) ToPointer() *InvalidParameterMaximumLengthRule {
	return &e
}
func (e *InvalidParameterMaximumLengthRule) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "max_length":
		fallthrough
	case "max_items":
		fallthrough
	case "max":
		*e = InvalidParameterMaximumLengthRule(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InvalidParameterMaximumLengthRule: %v", v)
	}
}

type InvalidParameterMaximumLength struct {
	Field string `json:"field"`
	// invalid parameters rules
	Rule    InvalidParameterMaximumLengthRule `json:"rule"`
	Maximum int64                             `json:"maximum"`
	Source  *string                           `json:"source,omitempty"`
	Reason  string                            `json:"reason"`
}

func (o *InvalidParameterMaximumLength) GetField() string {
	if o == nil {
		return ""
	}
	return o.Field
}

func (o *InvalidParameterMaximumLength) GetRule() InvalidParameterMaximumLengthRule {
	if o == nil {
		return InvalidParameterMaximumLengthRule("")
	}
	return o.Rule
}

func (o *InvalidParameterMaximumLength) GetMaximum() int64 {
	if o == nil {
		return 0
	}
	return o.Maximum
}

func (o *InvalidParameterMaximumLength) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *InvalidParameterMaximumLength) GetReason() string {
	if o == nil {
		return ""
	}
	return o.Reason
}
