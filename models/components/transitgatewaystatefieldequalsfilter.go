// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/utils"
)

type TransitGatewayStateFieldEqualsComparison struct {
}

type TransitGatewayStateFieldEqualsFilterType string

const (
	TransitGatewayStateFieldEqualsFilterTypeTransitGatewayState                      TransitGatewayStateFieldEqualsFilterType = "TransitGatewayState"
	TransitGatewayStateFieldEqualsFilterTypeTransitGatewayStateFieldEqualsComparison TransitGatewayStateFieldEqualsFilterType = "TransitGatewayStateFieldEqualsComparison"
)

// TransitGatewayStateFieldEqualsFilter - Filter transit-gateway state by exact match.
type TransitGatewayStateFieldEqualsFilter struct {
	TransitGatewayState                      *TransitGatewayState                      `queryParam:"inline"`
	TransitGatewayStateFieldEqualsComparison *TransitGatewayStateFieldEqualsComparison `queryParam:"inline"`

	Type TransitGatewayStateFieldEqualsFilterType
}

func CreateTransitGatewayStateFieldEqualsFilterTransitGatewayState(transitGatewayState TransitGatewayState) TransitGatewayStateFieldEqualsFilter {
	typ := TransitGatewayStateFieldEqualsFilterTypeTransitGatewayState

	return TransitGatewayStateFieldEqualsFilter{
		TransitGatewayState: &transitGatewayState,
		Type:                typ,
	}
}

func CreateTransitGatewayStateFieldEqualsFilterTransitGatewayStateFieldEqualsComparison(transitGatewayStateFieldEqualsComparison TransitGatewayStateFieldEqualsComparison) TransitGatewayStateFieldEqualsFilter {
	typ := TransitGatewayStateFieldEqualsFilterTypeTransitGatewayStateFieldEqualsComparison

	return TransitGatewayStateFieldEqualsFilter{
		TransitGatewayStateFieldEqualsComparison: &transitGatewayStateFieldEqualsComparison,
		Type:                                     typ,
	}
}

func (u *TransitGatewayStateFieldEqualsFilter) UnmarshalJSON(data []byte) error {

	var transitGatewayStateFieldEqualsComparison TransitGatewayStateFieldEqualsComparison = TransitGatewayStateFieldEqualsComparison{}
	if err := utils.UnmarshalJSON(data, &transitGatewayStateFieldEqualsComparison, "", true, true); err == nil {
		u.TransitGatewayStateFieldEqualsComparison = &transitGatewayStateFieldEqualsComparison
		u.Type = TransitGatewayStateFieldEqualsFilterTypeTransitGatewayStateFieldEqualsComparison
		return nil
	}

	var transitGatewayState TransitGatewayState = TransitGatewayState("")
	if err := utils.UnmarshalJSON(data, &transitGatewayState, "", true, true); err == nil {
		u.TransitGatewayState = &transitGatewayState
		u.Type = TransitGatewayStateFieldEqualsFilterTypeTransitGatewayState
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for TransitGatewayStateFieldEqualsFilter", string(data))
}

func (u TransitGatewayStateFieldEqualsFilter) MarshalJSON() ([]byte, error) {
	if u.TransitGatewayState != nil {
		return utils.MarshalJSON(u.TransitGatewayState, "", true)
	}

	if u.TransitGatewayStateFieldEqualsComparison != nil {
		return utils.MarshalJSON(u.TransitGatewayStateFieldEqualsComparison, "", true)
	}

	return nil, errors.New("could not marshal union type TransitGatewayStateFieldEqualsFilter: all fields are null")
}
