// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// UpdateTeam - The request schema for the update team request.
type UpdateTeam struct {
	// The name of the team.
	Name *string `json:"name,omitempty"`
	// The description of the team.
	Description *string `json:"description,omitempty"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	//
	// Labels are intended to store **INTERNAL** metadata.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	Labels map[string]*string `json:"labels,omitempty"`
}

func (o *UpdateTeam) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateTeam) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateTeam) GetLabels() map[string]*string {
	if o == nil {
		return nil
	}
	return o.Labels
}
