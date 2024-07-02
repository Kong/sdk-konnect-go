// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

// GroupConflictResource - A resource causing a conflict in a control plane group.
type GroupConflictResource struct {
	// The ID of the resource.
	ID string `json:"id"`
	// The type of the resource.
	Type string `json:"type"`
}

func (o *GroupConflictResource) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GroupConflictResource) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}