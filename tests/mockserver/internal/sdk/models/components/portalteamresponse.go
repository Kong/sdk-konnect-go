// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"mockserver/internal/sdk/utils"
	"time"
)

// PortalTeamResponse - Details about a developer team.
type PortalTeamResponse struct {
	ID          *string    `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func (p PortalTeamResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PortalTeamResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PortalTeamResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PortalTeamResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PortalTeamResponse) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PortalTeamResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PortalTeamResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
