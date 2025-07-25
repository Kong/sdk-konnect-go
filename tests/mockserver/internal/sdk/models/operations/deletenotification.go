// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type DeleteNotificationRequest struct {
	// ID of the notification.
	NotificationID string `pathParam:"style=simple,explode=false,name=notificationId"`
}

func (o *DeleteNotificationRequest) GetNotificationID() string {
	if o == nil {
		return ""
	}
	return o.NotificationID
}

type DeleteNotificationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteNotificationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
