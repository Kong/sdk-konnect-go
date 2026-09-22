package components

import "encoding/json"

// MarshalJSON serializes an UpdatePortalAuditLogWebhook without adding values for unset fields.
func (u UpdatePortalAuditLogWebhook) MarshalJSON() ([]byte, error) {
	type updatePortalAuditLogWebhookJSON UpdatePortalAuditLogWebhook
	return json.Marshal(updatePortalAuditLogWebhookJSON(u))
}

// UnmarshalJSON deserializes an UpdatePortalAuditLogWebhook without adding values for missing fields.
func (u *UpdatePortalAuditLogWebhook) UnmarshalJSON(data []byte) error {
	type updatePortalAuditLogWebhookJSON UpdatePortalAuditLogWebhook
	return json.Unmarshal(data, (*updatePortalAuditLogWebhookJSON)(u))
}
