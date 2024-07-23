// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

// VaultConfig - The configuration properties for the Vault which can be found on the vaults' documentation page.
type VaultConfig struct {
}

type Vault struct {
	// The configuration properties for the Vault which can be found on the vaults' documentation page.
	Config *VaultConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The description of the Vault entity.
	Description *string `json:"description,omitempty"`
	ID          *string `json:"id,omitempty"`
	// The name of the Vault that's going to be added. Currently, the Vault implementation must be installed in every Kong instance.
	Name *string `json:"name,omitempty"`
	// The unique prefix (or identifier) for this Vault configuration. The prefix is used to load the right Vault configuration and implementation when referencing secrets with the other entities.
	Prefix *string `json:"prefix,omitempty"`
	// An optional set of strings associated with the Vault for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (o *Vault) GetConfig() *VaultConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *Vault) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Vault) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Vault) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Vault) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Vault) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *Vault) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Vault) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

type VaultInput struct {
	// The configuration properties for the Vault which can be found on the vaults' documentation page.
	Config *VaultConfig `json:"config,omitempty"`
	// The description of the Vault entity.
	Description *string `json:"description,omitempty"`
	// The name of the Vault that's going to be added. Currently, the Vault implementation must be installed in every Kong instance.
	Name *string `json:"name,omitempty"`
	// The unique prefix (or identifier) for this Vault configuration. The prefix is used to load the right Vault configuration and implementation when referencing secrets with the other entities.
	Prefix *string `json:"prefix,omitempty"`
	// An optional set of strings associated with the Vault for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (o *VaultInput) GetConfig() *VaultConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *VaultInput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *VaultInput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *VaultInput) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *VaultInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
