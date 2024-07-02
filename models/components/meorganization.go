// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"time"
)

// MeOrganizationState - State of the organization
type MeOrganizationState string

const (
	MeOrganizationStateActive   MeOrganizationState = "active"
	MeOrganizationStateInactive MeOrganizationState = "inactive"
)

func (e MeOrganizationState) ToPointer() *MeOrganizationState {
	return &e
}
func (e *MeOrganizationState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "inactive":
		*e = MeOrganizationState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeOrganizationState: %v", v)
	}
}

type MeOrganization struct {
	// UUID of the organization.
	ID *string `json:"id,omitempty"`
	// Name of the organization.
	Name *string `json:"name,omitempty"`
	// Owner ID of the organization.
	OwnerID *string `json:"owner_id,omitempty"`
	// Path to organization-specific login when single sign on (SSO) is enabled. Blank otherwise.
	LoginPath *string `json:"login_path,omitempty"`
	// Date the organization was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date the organization was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// State of the organization
	State *MeOrganizationState `json:"state,omitempty"`
}

func (m MeOrganization) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeOrganization) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeOrganization) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *MeOrganization) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeOrganization) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *MeOrganization) GetLoginPath() *string {
	if o == nil {
		return nil
	}
	return o.LoginPath
}

func (o *MeOrganization) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *MeOrganization) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *MeOrganization) GetState() *MeOrganizationState {
	if o == nil {
		return nil
	}
	return o.State
}