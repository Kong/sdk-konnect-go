// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type CreateBasicAuthWithoutParents struct {
	Password *string  `json:"password,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Username *string  `json:"username,omitempty"`
}

func (o *CreateBasicAuthWithoutParents) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *CreateBasicAuthWithoutParents) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateBasicAuthWithoutParents) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}