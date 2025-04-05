// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type BasicAuthWithoutParentsConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *BasicAuthWithoutParentsConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type BasicAuthWithoutParents struct {
	Consumer *BasicAuthWithoutParentsConsumer `json:"consumer,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64   `json:"created_at,omitempty"`
	ID        *string  `json:"id,omitempty"`
	Password  string   `json:"password"`
	Tags      []string `json:"tags,omitempty"`
	Username  string   `json:"username"`
}

func (o *BasicAuthWithoutParents) GetConsumer() *BasicAuthWithoutParentsConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *BasicAuthWithoutParents) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *BasicAuthWithoutParents) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BasicAuthWithoutParents) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *BasicAuthWithoutParents) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *BasicAuthWithoutParents) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
