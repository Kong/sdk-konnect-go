// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"mockserver/internal/sdk/utils"
)

// CreateConfigurationRequest - Request schema for creating a configuration.
type CreateConfigurationRequest struct {
	ControlPlaneID string `json:"control_plane_id"`
	// Set of control-plane geos supported for deploying cloud-gateways configurations.
	ControlPlaneGeo ControlPlaneGeo `json:"control_plane_geo"`
	// Supported gateway version.
	Version string `json:"version"`
	// List of data-plane groups that describe where to deploy instances, along with how many instances.
	DataplaneGroups []CreateConfigurationDataPlaneGroup `json:"dataplane_groups"`
	// Type of API access data-plane groups will support for a configuration.
	APIAccess *APIAccess `default:"private+public" json:"api_access"`
}

func (c CreateConfigurationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateConfigurationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateConfigurationRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateConfigurationRequest) GetControlPlaneGeo() ControlPlaneGeo {
	if o == nil {
		return ControlPlaneGeo("")
	}
	return o.ControlPlaneGeo
}

func (o *CreateConfigurationRequest) GetVersion() string {
	if o == nil {
		return ""
	}
	return o.Version
}

func (o *CreateConfigurationRequest) GetDataplaneGroups() []CreateConfigurationDataPlaneGroup {
	if o == nil {
		return []CreateConfigurationDataPlaneGroup{}
	}
	return o.DataplaneGroups
}

func (o *CreateConfigurationRequest) GetAPIAccess() *APIAccess {
	if o == nil {
		return nil
	}
	return o.APIAccess
}
