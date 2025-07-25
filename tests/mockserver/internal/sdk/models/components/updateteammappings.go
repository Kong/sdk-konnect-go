// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Mapping struct {
	Group   *string  `json:"group,omitempty"`
	TeamIds []string `json:"team_ids,omitempty"`
}

func (o *Mapping) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *Mapping) GetTeamIds() []string {
	if o == nil {
		return nil
	}
	return o.TeamIds
}

// UpdateTeamMappings - The request schema for updating IdP team mappings.
type UpdateTeamMappings struct {
	// The mappings object.
	Mappings []Mapping `json:"mappings,omitempty"`
}

func (o *UpdateTeamMappings) GetMappings() []Mapping {
	if o == nil {
		return nil
	}
	return o.Mappings
}
