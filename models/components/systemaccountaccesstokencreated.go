// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"time"
)

type SystemAccountAccessTokenCreated struct {
	// ID of the system account access token.
	ID *string `json:"id,omitempty"`
	// Name of the system account access token.
	Name *string `json:"name,omitempty"`
	// Timestamp of when the system account access token was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp of when the system account access token was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Timestamp of when the system account access token will expire.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// Timestamp of when the system account access token was last used.
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
	// The token of the system account access token.
	Token *string `json:"token,omitempty"`
}

func (s SystemAccountAccessTokenCreated) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SystemAccountAccessTokenCreated) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SystemAccountAccessTokenCreated) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SystemAccountAccessTokenCreated) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SystemAccountAccessTokenCreated) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *SystemAccountAccessTokenCreated) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *SystemAccountAccessTokenCreated) GetExpiresAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *SystemAccountAccessTokenCreated) GetLastUsedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastUsedAt
}

func (o *SystemAccountAccessTokenCreated) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}
