// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"time"
)

// Details - Metadata associated with the notification to build actionable links.
type Details struct {
}

// Notification - Properties of a notification.
type Notification struct {
	// Contains a unique identifier used for this resource.
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// Status of the notification.
	Status NotificationStatus `json:"status"`
	// Region associated to a notification.
	Region    NotificationRegion    `json:"region"`
	Namespace NotificationNamespace `json:"namespace"`
	// ID of the entity. Use '*' to represent all entities of this type.
	EntityID string `json:"entity_id"`
	// Metadata associated with the notification to build actionable links.
	Details Details `json:"details"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt time.Time `json:"created_at"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (n Notification) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *Notification) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Notification) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Notification) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *Notification) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *Notification) GetStatus() NotificationStatus {
	if o == nil {
		return NotificationStatus("")
	}
	return o.Status
}

func (o *Notification) GetRegion() NotificationRegion {
	if o == nil {
		return NotificationRegion("")
	}
	return o.Region
}

func (o *Notification) GetNamespace() NotificationNamespace {
	if o == nil {
		return NotificationNamespace("")
	}
	return o.Namespace
}

func (o *Notification) GetEntityID() string {
	if o == nil {
		return ""
	}
	return o.EntityID
}

func (o *Notification) GetDetails() Details {
	if o == nil {
		return Details{}
	}
	return o.Details
}

func (o *Notification) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Notification) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
