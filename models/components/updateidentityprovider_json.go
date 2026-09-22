package components

import "encoding/json"

// MarshalJSON serializes an UpdateIdentityProvider without adding values for unset fields.
func (u UpdateIdentityProvider) MarshalJSON() ([]byte, error) {
	type updateIdentityProviderJSON UpdateIdentityProvider
	return json.Marshal(updateIdentityProviderJSON(u))
}

// UnmarshalJSON deserializes an UpdateIdentityProvider without adding values for missing fields.
func (u *UpdateIdentityProvider) UnmarshalJSON(data []byte) error {
	type updateIdentityProviderJSON UpdateIdentityProvider
	return json.Unmarshal(data, (*updateIdentityProviderJSON)(u))
}
