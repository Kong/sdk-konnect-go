// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// TeamMapping - A team assignment is a mapping of an IdP group to a Konnect Team.
type TeamMapping struct {
	// The IdP group.
	Group *string `json:"group,omitempty"`
	// An array of ID's that are mapped to the specified group.
	TeamIds []string `json:"team_ids,omitempty"`
}

func (o *TeamMapping) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *TeamMapping) GetTeamIds() []string {
	if o == nil {
		return nil
	}
	return o.TeamIds
}
