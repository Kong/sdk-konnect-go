// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type CreateHMACAuthWithoutParents struct {
	Secret   *string  `json:"secret,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Username *string  `json:"username,omitempty"`
}

func (o *CreateHMACAuthWithoutParents) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *CreateHMACAuthWithoutParents) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateHMACAuthWithoutParents) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}
