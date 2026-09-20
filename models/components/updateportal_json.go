package components

import "encoding/json"

// MarshalJSON serializes an UpdatePortal without adding values for unset fields.
func (u UpdatePortal) MarshalJSON() ([]byte, error) {
	type updatePortalJSON UpdatePortal
	return json.Marshal(updatePortalJSON(u))
}

// UnmarshalJSON deserializes an UpdatePortal without adding values for missing fields.
func (u *UpdatePortal) UnmarshalJSON(data []byte) error {
	type updatePortalJSON UpdatePortal
	return json.Unmarshal(data, (*updatePortalJSON)(u))
}
