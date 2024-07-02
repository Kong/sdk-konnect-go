// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// AssignRoleEntityRegion - The region of the team.
type AssignRoleEntityRegion string

const (
	AssignRoleEntityRegionUs       AssignRoleEntityRegion = "us"
	AssignRoleEntityRegionEu       AssignRoleEntityRegion = "eu"
	AssignRoleEntityRegionAu       AssignRoleEntityRegion = "au"
	AssignRoleEntityRegionWildcard AssignRoleEntityRegion = "*"
)

func (e AssignRoleEntityRegion) ToPointer() *AssignRoleEntityRegion {
	return &e
}
func (e *AssignRoleEntityRegion) UnmarshalJSON(data []byte) error {
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
	case "*":
		*e = AssignRoleEntityRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssignRoleEntityRegion: %v", v)
	}
}

// AssignRole - An assigned role is a role that has been assigned to a user or team.
type AssignRole struct {
	// The desired role.
	RoleName *string `json:"role_name,omitempty"`
	// The ID of the entity.
	EntityID *string `json:"entity_id,omitempty"`
	// The type of entity.
	EntityTypeName *string `json:"entity_type_name,omitempty"`
	// The region of the team.
	EntityRegion *AssignRoleEntityRegion `json:"entity_region,omitempty"`
}

func (o *AssignRole) GetRoleName() *string {
	if o == nil {
		return nil
	}
	return o.RoleName
}

func (o *AssignRole) GetEntityID() *string {
	if o == nil {
		return nil
	}
	return o.EntityID
}

func (o *AssignRole) GetEntityTypeName() *string {
	if o == nil {
		return nil
	}
	return o.EntityTypeName
}

func (o *AssignRole) GetEntityRegion() *AssignRoleEntityRegion {
	if o == nil {
		return nil
	}
	return o.EntityRegion
}
