// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetPortalAuditLogWebhookStatusRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
}

func (o *GetPortalAuditLogWebhookStatusRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

type GetPortalAuditLogWebhookStatusResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Get response for portal audit log webhook status
	PortalAuditLogWebhookStatus *components.PortalAuditLogWebhookStatus
}

func (o *GetPortalAuditLogWebhookStatusResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPortalAuditLogWebhookStatusResponse) GetPortalAuditLogWebhookStatus() *components.PortalAuditLogWebhookStatus {
	if o == nil {
		return nil
	}
	return o.PortalAuditLogWebhookStatus
}
