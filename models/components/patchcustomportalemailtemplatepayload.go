// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// PatchCustomPortalEmailTemplatePayload - Update a custom portal email template.
type PatchCustomPortalEmailTemplatePayload struct {
	Content *EmailTemplateContent `json:"content,omitempty"`
	// Whether the email template is enabled or disabled for a portal
	Enabled *bool `json:"enabled,omitempty"`
}

func (o *PatchCustomPortalEmailTemplatePayload) GetContent() *EmailTemplateContent {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *PatchCustomPortalEmailTemplatePayload) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}
