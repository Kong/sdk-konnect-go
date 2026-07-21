package components

import "encoding/json"

// MarshalJSON serializes a PortalUpdateIdentityProvider without adding values for unset fields.
func (p PortalUpdateIdentityProvider) MarshalJSON() ([]byte, error) {
	type portalUpdateIdentityProviderJSON PortalUpdateIdentityProvider
	return json.Marshal(portalUpdateIdentityProviderJSON(p))
}

// UnmarshalJSON deserializes a PortalUpdateIdentityProvider without adding values for missing fields.
func (p *PortalUpdateIdentityProvider) UnmarshalJSON(data []byte) error {
	type portalUpdateIdentityProviderJSON PortalUpdateIdentityProvider
	return json.Unmarshal(data, (*portalUpdateIdentityProviderJSON)(p))
}
