// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// EntityRegion - Region of the entity.
type EntityRegion string

const (
	EntityRegionUs       EntityRegion = "us"
	EntityRegionEu       EntityRegion = "eu"
	EntityRegionAu       EntityRegion = "au"
	EntityRegionMe       EntityRegion = "me"
	EntityRegionIn       EntityRegion = "in"
	EntityRegionWildcard EntityRegion = "*"
)

func (e EntityRegion) ToPointer() *EntityRegion {
	return &e
}
func (e *EntityRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "us":
		fallthrough
	case "eu":
		fallthrough
	case "au":
		fallthrough
	case "me":
		fallthrough
	case "in":
		fallthrough
	case "*":
		*e = EntityRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EntityRegion: %v", v)
	}
}

// PortalAssignedRoleResponse - An assigned role associates a service and an action to a team.
type PortalAssignedRoleResponse struct {
	ID             *string `json:"id,omitempty"`
	RoleName       *string `json:"role_name,omitempty"`
	EntityID       *string `json:"entity_id,omitempty"`
	EntityTypeName *string `json:"entity_type_name,omitempty"`
	// Region of the entity.
	EntityRegion *EntityRegion `json:"entity_region,omitempty"`
}

func (o *PortalAssignedRoleResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PortalAssignedRoleResponse) GetRoleName() *string {
	if o == nil {
		return nil
	}
	return o.RoleName
}

func (o *PortalAssignedRoleResponse) GetEntityID() *string {
	if o == nil {
		return nil
	}
	return o.EntityID
}

func (o *PortalAssignedRoleResponse) GetEntityTypeName() *string {
	if o == nil {
		return nil
	}
	return o.EntityTypeName
}

func (o *PortalAssignedRoleResponse) GetEntityRegion() *EntityRegion {
	if o == nil {
		return nil
	}
	return o.EntityRegion
}
