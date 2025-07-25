// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ListEmailTemplates - A list of customized email templates used in a portal.
type ListEmailTemplates struct {
	Data []EmailTemplate `json:"data,omitempty"`
}

func (o *ListEmailTemplates) GetData() []EmailTemplate {
	if o == nil {
		return nil
	}
	return o.Data
}
