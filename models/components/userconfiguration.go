// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// UserConfiguration - Properties of an event.
type UserConfiguration struct {
	EventID                string                `json:"event_id"`
	EventTitle             string                `json:"event_title"`
	EventDescription       string                `json:"event_description"`
	EventNamespace         NotificationNamespace `json:"event_namespace"`
	EventSubscriptionCount int64                 `json:"event_subscription_count"`
	DefaultSubscription    []DefaultSubscription `json:"default_subscription"`
}

func (o *UserConfiguration) GetEventID() string {
	if o == nil {
		return ""
	}
	return o.EventID
}

func (o *UserConfiguration) GetEventTitle() string {
	if o == nil {
		return ""
	}
	return o.EventTitle
}

func (o *UserConfiguration) GetEventDescription() string {
	if o == nil {
		return ""
	}
	return o.EventDescription
}

func (o *UserConfiguration) GetEventNamespace() NotificationNamespace {
	if o == nil {
		return NotificationNamespace("")
	}
	return o.EventNamespace
}

func (o *UserConfiguration) GetEventSubscriptionCount() int64 {
	if o == nil {
		return 0
	}
	return o.EventSubscriptionCount
}

func (o *UserConfiguration) GetDefaultSubscription() []DefaultSubscription {
	if o == nil {
		return []DefaultSubscription{}
	}
	return o.DefaultSubscription
}
