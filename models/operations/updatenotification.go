// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type UpdateNotificationRequest struct {
	// ID of the notification.
	NotificationID string `pathParam:"style=simple,explode=false,name=notificationId"`
	// Request body schema for updating notification status.
	NotificationUpdatePayload *components.NotificationUpdatePayload `request:"mediaType=application/json"`
}

func (o *UpdateNotificationRequest) GetNotificationID() string {
	if o == nil {
		return ""
	}
	return o.NotificationID
}

func (o *UpdateNotificationRequest) GetNotificationUpdatePayload() *components.NotificationUpdatePayload {
	if o == nil {
		return nil
	}
	return o.NotificationUpdatePayload
}

type UpdateNotificationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// View a single notification.
	Notification *components.Notification
}

func (o *UpdateNotificationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateNotificationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateNotificationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateNotificationResponse) GetNotification() *components.Notification {
	if o == nil {
		return nil
	}
	return o.Notification
}
