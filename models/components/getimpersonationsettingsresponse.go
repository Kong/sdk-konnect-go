// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

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