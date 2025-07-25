// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"time"
)

// BasicDeveloper - Basic information about a developer.
type BasicDeveloper struct {
	ID        *string    `json:"id,omitempty"`
	Email     *string    `json:"email,omitempty"`
	FullName  *string    `json:"full_name,omitempty"`
	Active    *bool      `json:"active,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (b BasicDeveloper) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BasicDeveloper) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BasicDeveloper) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BasicDeveloper) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *BasicDeveloper) GetFullName() *string {
	if o == nil {
		return nil
	}
	return o.FullName
}

func (o *BasicDeveloper) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *BasicDeveloper) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *BasicDeveloper) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
