package components

import "encoding/json"

// MarshalJSON serializes a PortalUpdateTeamRequest without adding unset fields.
func (p PortalUpdateTeamRequest) MarshalJSON() ([]byte, error) {
	type portalUpdateTeamJSON PortalUpdateTeamRequest
	return json.Marshal(portalUpdateTeamJSON(p))
}

// UnmarshalJSON decodes a PortalUpdateTeamRequest without adding missing fields.
func (p *PortalUpdateTeamRequest) UnmarshalJSON(data []byte) error {
	type portalUpdateTeamJSON PortalUpdateTeamRequest
	return json.Unmarshal(data, (*portalUpdateTeamJSON)(p))
}
