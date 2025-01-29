// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"time"
)

// State of the data-plane group.
type State string

const (
	StateCreated      State = "created"
	StateInitializing State = "initializing"
	StateReady        State = "ready"
	StateTerminating  State = "terminating"
	StateTerminated   State = "terminated"
)

func (e State) ToPointer() *State {
	return &e
}
func (e *State) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created":
		fallthrough
	case "initializing":
		fallthrough
	case "ready":
		fallthrough
	case "terminating":
		fallthrough
	case "terminated":
		*e = State(v)
		return nil
	default:
		return fmt.Errorf("invalid value for State: %v", v)
	}
}

// ConfigurationDataPlaneGroup - Object that describes the set of data-plane groups currently pointed to this configuration.
type ConfigurationDataPlaneGroup struct {
	// ID of the data-plane group that represents a deployment target for a set of data-planes.
	ID string `json:"id"`
	// Name of cloud provider.
	Provider ProviderName `json:"provider"`
	// Region ID for cloud provider region.
	Region    string                               `json:"region"`
	Autoscale ConfigurationDataPlaneGroupAutoscale `json:"autoscale"`
	// Array of environment variables to set for a data-plane group.
	Environment           []ConfigurationDataPlaneGroupEnvironmentField `json:"environment,omitempty"`
	CloudGatewayNetworkID string                                        `json:"cloud_gateway_network_id"`
	// State of the data-plane group.
	State State `json:"state"`
	// List of private IP addresses of the internal load balancer that proxies traffic to this data-plane group.
	//
	PrivateIPAddresses []string `json:"private_ip_addresses,omitempty"`
	// List of egress IP addresses for the network that this data-plane group runs on.
	//
	EgressIPAddresses []string `json:"egress_ip_addresses,omitempty"`
	// An RFC-3339 timestamp representation of data-plane group creation date.
	CreatedAt time.Time `json:"created_at"`
	// An RFC-3339 timestamp representation of data-plane group update date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (c ConfigurationDataPlaneGroup) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConfigurationDataPlaneGroup) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConfigurationDataPlaneGroup) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConfigurationDataPlaneGroup) GetProvider() ProviderName {
	if o == nil {
		return ProviderName("")
	}
	return o.Provider
}

func (o *ConfigurationDataPlaneGroup) GetRegion() string {
	if o == nil {
		return ""
	}
	return o.Region
}

func (o *ConfigurationDataPlaneGroup) GetAutoscale() ConfigurationDataPlaneGroupAutoscale {
	if o == nil {
		return ConfigurationDataPlaneGroupAutoscale{}
	}
	return o.Autoscale
}

func (o *ConfigurationDataPlaneGroup) GetEnvironment() []ConfigurationDataPlaneGroupEnvironmentField {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *ConfigurationDataPlaneGroup) GetCloudGatewayNetworkID() string {
	if o == nil {
		return ""
	}
	return o.CloudGatewayNetworkID
}

func (o *ConfigurationDataPlaneGroup) GetState() State {
	if o == nil {
		return State("")
	}
	return o.State
}

func (o *ConfigurationDataPlaneGroup) GetPrivateIPAddresses() []string {
	if o == nil {
		return nil
	}
	return o.PrivateIPAddresses
}

func (o *ConfigurationDataPlaneGroup) GetEgressIPAddresses() []string {
	if o == nil {
		return nil
	}
	return o.EgressIPAddresses
}

func (o *ConfigurationDataPlaneGroup) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ConfigurationDataPlaneGroup) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
