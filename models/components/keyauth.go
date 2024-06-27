// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type KeyAuthConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *KeyAuthConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type KeyAuth struct {
	Key      *string          `json:"key,omitempty"`
	Tags     []string         `json:"tags,omitempty"`
	Consumer *KeyAuthConsumer `json:"consumer,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64  `json:"created_at,omitempty"`
	ID        *string `json:"id,omitempty"`
}

func (o *KeyAuth) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *KeyAuth) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *KeyAuth) GetConsumer() *KeyAuthConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *KeyAuth) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *KeyAuth) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
