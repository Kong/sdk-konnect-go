package components

import "encoding/json"

// MarshalJSON serializes a PatchCustomPortalEmailTemplatePayload without adding values for unset fields.
func (p PatchCustomPortalEmailTemplatePayload) MarshalJSON() ([]byte, error) {
	type patchCustomPortalEmailTemplatePayloadJSON PatchCustomPortalEmailTemplatePayload
	return json.Marshal(patchCustomPortalEmailTemplatePayloadJSON(p))
}

// UnmarshalJSON deserializes a PatchCustomPortalEmailTemplatePayload without adding values for missing fields.
func (p *PatchCustomPortalEmailTemplatePayload) UnmarshalJSON(data []byte) error {
	type patchCustomPortalEmailTemplatePayloadJSON PatchCustomPortalEmailTemplatePayload
	return json.Unmarshal(data, (*patchCustomPortalEmailTemplatePayloadJSON)(p))
}
