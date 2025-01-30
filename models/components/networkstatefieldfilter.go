// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/utils"
)

type NetworkStateFieldFilterType string

const (
	NetworkStateFieldFilterTypeNetworkStateFieldEqualsFilter     NetworkStateFieldFilterType = "NetworkStateFieldEqualsFilter"
	NetworkStateFieldFilterTypeNetworkStateFieldNotEqualsFilter  NetworkStateFieldFilterType = "NetworkStateFieldNotEqualsFilter"
	NetworkStateFieldFilterTypeNetworkStateFieldOrEqualityFilter NetworkStateFieldFilterType = "NetworkStateFieldOrEqualityFilter"
)

type NetworkStateFieldFilter struct {
	NetworkStateFieldEqualsFilter     *NetworkStateFieldEqualsFilter     `queryParam:"inline"`
	NetworkStateFieldNotEqualsFilter  *NetworkStateFieldNotEqualsFilter  `queryParam:"inline"`
	NetworkStateFieldOrEqualityFilter *NetworkStateFieldOrEqualityFilter `queryParam:"inline"`

	Type NetworkStateFieldFilterType
}

func CreateNetworkStateFieldFilterNetworkStateFieldEqualsFilter(networkStateFieldEqualsFilter NetworkStateFieldEqualsFilter) NetworkStateFieldFilter {
	typ := NetworkStateFieldFilterTypeNetworkStateFieldEqualsFilter

	return NetworkStateFieldFilter{
		NetworkStateFieldEqualsFilter: &networkStateFieldEqualsFilter,
		Type:                          typ,
	}
}

func CreateNetworkStateFieldFilterNetworkStateFieldNotEqualsFilter(networkStateFieldNotEqualsFilter NetworkStateFieldNotEqualsFilter) NetworkStateFieldFilter {
	typ := NetworkStateFieldFilterTypeNetworkStateFieldNotEqualsFilter

	return NetworkStateFieldFilter{
		NetworkStateFieldNotEqualsFilter: &networkStateFieldNotEqualsFilter,
		Type:                             typ,
	}
}

func CreateNetworkStateFieldFilterNetworkStateFieldOrEqualityFilter(networkStateFieldOrEqualityFilter NetworkStateFieldOrEqualityFilter) NetworkStateFieldFilter {
	typ := NetworkStateFieldFilterTypeNetworkStateFieldOrEqualityFilter

	return NetworkStateFieldFilter{
		NetworkStateFieldOrEqualityFilter: &networkStateFieldOrEqualityFilter,
		Type:                              typ,
	}
}

func (u *NetworkStateFieldFilter) UnmarshalJSON(data []byte) error {

	var networkStateFieldNotEqualsFilter NetworkStateFieldNotEqualsFilter = NetworkStateFieldNotEqualsFilter{}
	if err := utils.UnmarshalJSON(data, &networkStateFieldNotEqualsFilter, "", true, true); err == nil {
		u.NetworkStateFieldNotEqualsFilter = &networkStateFieldNotEqualsFilter
		u.Type = NetworkStateFieldFilterTypeNetworkStateFieldNotEqualsFilter
		return nil
	}

	var networkStateFieldOrEqualityFilter NetworkStateFieldOrEqualityFilter = NetworkStateFieldOrEqualityFilter{}
	if err := utils.UnmarshalJSON(data, &networkStateFieldOrEqualityFilter, "", true, true); err == nil {
		u.NetworkStateFieldOrEqualityFilter = &networkStateFieldOrEqualityFilter
		u.Type = NetworkStateFieldFilterTypeNetworkStateFieldOrEqualityFilter
		return nil
	}

	var networkStateFieldEqualsFilter NetworkStateFieldEqualsFilter = NetworkStateFieldEqualsFilter{}
	if err := utils.UnmarshalJSON(data, &networkStateFieldEqualsFilter, "", true, true); err == nil {
		u.NetworkStateFieldEqualsFilter = &networkStateFieldEqualsFilter
		u.Type = NetworkStateFieldFilterTypeNetworkStateFieldEqualsFilter
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for NetworkStateFieldFilter", string(data))
}

func (u NetworkStateFieldFilter) MarshalJSON() ([]byte, error) {
	if u.NetworkStateFieldEqualsFilter != nil {
		return utils.MarshalJSON(u.NetworkStateFieldEqualsFilter, "", true)
	}

	if u.NetworkStateFieldNotEqualsFilter != nil {
		return utils.MarshalJSON(u.NetworkStateFieldNotEqualsFilter, "", true)
	}

	if u.NetworkStateFieldOrEqualityFilter != nil {
		return utils.MarshalJSON(u.NetworkStateFieldOrEqualityFilter, "", true)
	}

	return nil, errors.New("could not marshal union type NetworkStateFieldFilter: all fields are null")
}
