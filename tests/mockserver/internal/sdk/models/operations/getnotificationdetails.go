// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetNotificationDetailsRequest struct {
	// ID of the notification.
	NotificationID string `pathParam:"style=simple,explode=false,name=notificationId"`
}

func (o *GetNotificationDetailsRequest) GetNotificationID() string {
	if o == nil {
		return ""
	}
	return o.NotificationID
}

type GetNotificationDetailsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// View a single notification.
	Notification *components.Notification
}

func (o *GetNotificationDetailsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetNotificationDetailsResponse) GetNotification() *components.Notification {
	if o == nil {
		return nil
	}
	return o.Notification
}
