// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type BasicAuthConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *BasicAuthConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type BasicAuth struct {
	Password *string            `json:"password,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
	Username *string            `json:"username,omitempty"`
	Consumer *BasicAuthConsumer `json:"consumer,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64  `json:"created_at,omitempty"`
	ID        *string `json:"id,omitempty"`
}

func (o *BasicAuth) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *BasicAuth) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *BasicAuth) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *BasicAuth) GetConsumer() *BasicAuthConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *BasicAuth) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *BasicAuth) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
