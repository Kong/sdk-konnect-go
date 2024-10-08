// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetImpersonationSettingsResponse - Response for Get Impersonation Settings endpoint
type GetImpersonationSettingsResponse struct {
	// The organization has user impersonation enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

func (o *GetImpersonationSettingsResponse) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}
