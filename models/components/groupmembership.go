// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Members struct {
	ID *string `json:"id,omitempty"`
}

func (o *Members) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// GroupMembership - Request body for upserting a list of child control planes to a control plane group membership.
type GroupMembership struct {
	Members []Members `json:"members,omitempty"`
}

func (o *GroupMembership) GetMembers() []Members {
	if o == nil {
		return nil
	}
	return o.Members
}
