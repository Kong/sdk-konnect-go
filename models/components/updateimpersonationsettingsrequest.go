// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type UpdateImpersonationSettingsRequest struct {
	// Indicates if user impersonation is allowed for the organization.
	Enabled *bool `json:"enabled,omitempty"`
}

func (o *UpdateImpersonationSettingsRequest) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}
