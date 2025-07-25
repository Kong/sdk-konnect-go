// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"time"
)

// EventSubscriptionResponse - Properties associated with an event subscription.
type EventSubscriptionResponse struct {
	// Contains a unique identifier used for this resource.
	ID         string      `json:"id"`
	EntityType EntityTypes `json:"entity_type"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt time.Time `json:"created_at"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt time.Time             `json:"updated_at"`
	Regions   []NotificationRegion  `json:"regions"`
	Entities  []string              `json:"entities"`
	Channels  []NotificationChannel `json:"channels"`
	// Enable/Disable complete subscription within an event.
	Enabled bool `json:"enabled"`
	// User provided name of the subscription.
	Name string `json:"name"`
}

func (e EventSubscriptionResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EventSubscriptionResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *EventSubscriptionResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *EventSubscriptionResponse) GetEntityType() EntityTypes {
	if o == nil {
		return EntityTypes("")
	}
	return o.EntityType
}

func (o *EventSubscriptionResponse) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *EventSubscriptionResponse) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *EventSubscriptionResponse) GetRegions() []NotificationRegion {
	if o == nil {
		return []NotificationRegion{}
	}
	return o.Regions
}

func (o *EventSubscriptionResponse) GetEntities() []string {
	if o == nil {
		return []string{}
	}
	return o.Entities
}

func (o *EventSubscriptionResponse) GetChannels() []NotificationChannel {
	if o == nil {
		return []NotificationChannel{}
	}
	return o.Channels
}

func (o *EventSubscriptionResponse) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *EventSubscriptionResponse) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
