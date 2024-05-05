// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

// GroupMemberStatus - Object with information determining the group membership status of a control plane.
type GroupMemberStatus struct {
	// Boolean indicating if a control plane is a member of a control plane group.
	IsMember bool `json:"is_member"`
}

func (o *GroupMemberStatus) GetIsMember() bool {
	if o == nil {
		return false
	}
	return o.IsMember
}
