// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type ListNotificationsRequest struct {
	// Request the next page of data, starting with the item before this parameter.
	PageBefore *string `queryParam:"style=form,explode=true,name=page[before]"`
	// Request the next page of data, starting with the item after this parameter.
	PageAfter *string `queryParam:"style=form,explode=true,name=page[after]"`
	// Filters a collection of notifications.
	Filter *components.NotificationFilterParameters `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (o *ListNotificationsRequest) GetPageBefore() *string {
	if o == nil {
		return nil
	}
	return o.PageBefore
}

func (o *ListNotificationsRequest) GetPageAfter() *string {
	if o == nil {
		return nil
	}
	return o.PageAfter
}

func (o *ListNotificationsRequest) GetFilter() *components.NotificationFilterParameters {
	if o == nil {
		return nil
	}
	return o.Filter
}

type ListNotificationsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A paginated list response for a collection of notifications
	NotificationListResponse *components.NotificationListResponse
}

func (o *ListNotificationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListNotificationsResponse) GetNotificationListResponse() *components.NotificationListResponse {
	if o == nil {
		return nil
	}
	return o.NotificationListResponse
}
