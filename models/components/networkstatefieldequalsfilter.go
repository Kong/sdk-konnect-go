// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/utils"
)

type NetworkStateFieldEqualsComparison struct {
}

type NetworkStateFieldEqualsFilterType string

const (
	NetworkStateFieldEqualsFilterTypeNetworkState                      NetworkStateFieldEqualsFilterType = "NetworkState"
	NetworkStateFieldEqualsFilterTypeNetworkStateFieldEqualsComparison NetworkStateFieldEqualsFilterType = "NetworkStateFieldEqualsComparison"
)

// NetworkStateFieldEqualsFilter - Filter a network state by exact match.
type NetworkStateFieldEqualsFilter struct {
	NetworkState                      *NetworkState                      `queryParam:"inline"`
	NetworkStateFieldEqualsComparison *NetworkStateFieldEqualsComparison `queryParam:"inline"`

	Type NetworkStateFieldEqualsFilterType
}

func CreateNetworkStateFieldEqualsFilterNetworkState(networkState NetworkState) NetworkStateFieldEqualsFilter {
	typ := NetworkStateFieldEqualsFilterTypeNetworkState

	return NetworkStateFieldEqualsFilter{
		NetworkState: &networkState,
		Type:         typ,
	}
}

func CreateNetworkStateFieldEqualsFilterNetworkStateFieldEqualsComparison(networkStateFieldEqualsComparison NetworkStateFieldEqualsComparison) NetworkStateFieldEqualsFilter {
	typ := NetworkStateFieldEqualsFilterTypeNetworkStateFieldEqualsComparison

	return NetworkStateFieldEqualsFilter{
		NetworkStateFieldEqualsComparison: &networkStateFieldEqualsComparison,
		Type:                              typ,
	}
}

func (u *NetworkStateFieldEqualsFilter) UnmarshalJSON(data []byte) error {

	var networkStateFieldEqualsComparison NetworkStateFieldEqualsComparison = NetworkStateFieldEqualsComparison{}
	if err := utils.UnmarshalJSON(data, &networkStateFieldEqualsComparison, "", true, true); err == nil {
		u.NetworkStateFieldEqualsComparison = &networkStateFieldEqualsComparison
		u.Type = NetworkStateFieldEqualsFilterTypeNetworkStateFieldEqualsComparison
		return nil
	}

	var networkState NetworkState = NetworkState("")
	if err := utils.UnmarshalJSON(data, &networkState, "", true, true); err == nil {
		u.NetworkState = &networkState
		u.Type = NetworkStateFieldEqualsFilterTypeNetworkState
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for NetworkStateFieldEqualsFilter", string(data))
}

func (u NetworkStateFieldEqualsFilter) MarshalJSON() ([]byte, error) {
	if u.NetworkState != nil {
		return utils.MarshalJSON(u.NetworkState, "", true)
	}

	if u.NetworkStateFieldEqualsComparison != nil {
		return utils.MarshalJSON(u.NetworkStateFieldEqualsComparison, "", true)
	}

	return nil, errors.New("could not marshal union type NetworkStateFieldEqualsFilter: all fields are null")
}
