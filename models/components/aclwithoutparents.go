// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ACLWithoutParents struct {
	Group *string  `json:"group,omitempty"`
	ID    *string  `json:"id,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

func (o *ACLWithoutParents) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *ACLWithoutParents) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ACLWithoutParents) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
