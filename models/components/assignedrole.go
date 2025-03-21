// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// AssignedRole - An assigned role is a role that has been assigned to a user or team.
type AssignedRole struct {
	// The ID of the role assignment.
	ID *string `json:"id,omitempty"`
	// Name of the role being assigned.
	RoleName *string `json:"role_name,omitempty"`
	// A RBAC entity ID.
	EntityID *string `json:"entity_id,omitempty"`
	// Name of the entity type the role is being assigned to.
	EntityTypeName *string       `json:"entity_type_name,omitempty"`
	EntityRegion   *EntityRegion `json:"entity_region,omitempty"`
}

func (o *AssignedRole) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AssignedRole) GetRoleName() *string {
	if o == nil {
		return nil
	}
	return o.RoleName
}

func (o *AssignedRole) GetEntityID() *string {
	if o == nil {
		return nil
	}
	return o.EntityID
}

func (o *AssignedRole) GetEntityTypeName() *string {
	if o == nil {
		return nil
	}
	return o.EntityTypeName
}

func (o *AssignedRole) GetEntityRegion() *EntityRegion {
	if o == nil {
		return nil
	}
	return o.EntityRegion
}
