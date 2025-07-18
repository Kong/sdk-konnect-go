// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/utils"
	"time"
)

// GroupStatusState - The state of the control plane group.
type GroupStatusState string

const (
	GroupStatusStateOk       GroupStatusState = "OK"
	GroupStatusStateConflict GroupStatusState = "CONFLICT"
	GroupStatusStateUnknown  GroupStatusState = "UNKNOWN"
)

func (e GroupStatusState) ToPointer() *GroupStatusState {
	return &e
}
func (e *GroupStatusState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OK":
		fallthrough
	case "CONFLICT":
		fallthrough
	case "UNKNOWN":
		*e = GroupStatusState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GroupStatusState: %v", v)
	}
}

// GroupStatus - The Group Status object contains information about the status of a control plane group.
type GroupStatus struct {
	// The control plane group ID.
	ID string `json:"id"`
	// An ISO-8604 timestamp representation of control plane group status creation date.
	CreatedAt time.Time `json:"created_at"`
	// An ISO-8604 timestamp representation of control plane group status update date.
	UpdatedAt time.Time       `json:"updated_at"`
	Conflicts []GroupConflict `json:"conflicts,omitempty"`
	// The state of the control plane group.
	State GroupStatusState `json:"state"`
}

func (g GroupStatus) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GroupStatus) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GroupStatus) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GroupStatus) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *GroupStatus) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *GroupStatus) GetConflicts() []GroupConflict {
	if o == nil {
		return nil
	}
	return o.Conflicts
}

func (o *GroupStatus) GetState() GroupStatusState {
	if o == nil {
		return GroupStatusState("")
	}
	return o.State
}
