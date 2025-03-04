// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// KeyPem - A keypair in PEM format.
type KeyPem struct {
	PrivateKey *string `json:"private_key,omitempty"`
	PublicKey  *string `json:"public_key,omitempty"`
}

func (o *KeyPem) GetPrivateKey() *string {
	if o == nil {
		return nil
	}
	return o.PrivateKey
}

func (o *KeyPem) GetPublicKey() *string {
	if o == nil {
		return nil
	}
	return o.PublicKey
}

// KeySet1 - The id (an UUID) of the key-set with which to associate the key.
type KeySet1 struct {
	ID *string `json:"id,omitempty"`
}

func (o *KeySet1) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// A Key object holds a representation of asymmetric keys in various formats. When Kong or a Kong plugin requires a specific public or private key to perform certain operations, it can use this entity.
type Key struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64  `json:"created_at,omitempty"`
	ID        *string `json:"id,omitempty"`
	// A JSON Web Key represented as a string.
	Jwk *string `json:"jwk,omitempty"`
	// A unique identifier for a key.
	Kid string `json:"kid"`
	// The name to associate with the given keys.
	Name *string `json:"name,omitempty"`
	// A keypair in PEM format.
	Pem *KeyPem `json:"pem,omitempty"`
	// The id (an UUID) of the key-set with which to associate the key.
	Set *KeySet1 `json:"set,omitempty"`
	// An optional set of strings associated with the Key for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (o *Key) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Key) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Key) GetJwk() *string {
	if o == nil {
		return nil
	}
	return o.Jwk
}

func (o *Key) GetKid() string {
	if o == nil {
		return ""
	}
	return o.Kid
}

func (o *Key) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Key) GetPem() *KeyPem {
	if o == nil {
		return nil
	}
	return o.Pem
}

func (o *Key) GetSet() *KeySet1 {
	if o == nil {
		return nil
	}
	return o.Set
}

func (o *Key) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Key) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// KeyInput - A Key object holds a representation of asymmetric keys in various formats. When Kong or a Kong plugin requires a specific public or private key to perform certain operations, it can use this entity.
type KeyInput struct {
	ID *string `json:"id,omitempty"`
	// A JSON Web Key represented as a string.
	Jwk *string `json:"jwk,omitempty"`
	// A unique identifier for a key.
	Kid string `json:"kid"`
	// The name to associate with the given keys.
	Name *string `json:"name,omitempty"`
	// A keypair in PEM format.
	Pem *KeyPem `json:"pem,omitempty"`
	// The id (an UUID) of the key-set with which to associate the key.
	Set *KeySet1 `json:"set,omitempty"`
	// An optional set of strings associated with the Key for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (o *KeyInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *KeyInput) GetJwk() *string {
	if o == nil {
		return nil
	}
	return o.Jwk
}

func (o *KeyInput) GetKid() string {
	if o == nil {
		return ""
	}
	return o.Kid
}

func (o *KeyInput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *KeyInput) GetPem() *KeyPem {
	if o == nil {
		return nil
	}
	return o.Pem
}

func (o *KeyInput) GetSet() *KeySet1 {
	if o == nil {
		return nil
	}
	return o.Set
}

func (o *KeyInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
