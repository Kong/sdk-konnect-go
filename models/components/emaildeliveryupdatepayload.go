// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type EmailDeliveryUpdatePayload struct {
	Enabled      *bool   `json:"enabled,omitempty"`
	FromEmail    *string `json:"from_email,omitempty"`
	ReplyToEmail *string `json:"reply_to_email,omitempty"`
}

func (o *EmailDeliveryUpdatePayload) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *EmailDeliveryUpdatePayload) GetFromEmail() *string {
	if o == nil {
		return nil
	}
	return o.FromEmail
}

func (o *EmailDeliveryUpdatePayload) GetReplyToEmail() *string {
	if o == nil {
		return nil
	}
	return o.ReplyToEmail
}
