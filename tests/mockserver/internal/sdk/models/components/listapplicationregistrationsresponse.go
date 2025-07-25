// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ListApplicationRegistrationsResponse - A paginated list of application registrations.
type ListApplicationRegistrationsResponse struct {
	// returns the pagination information
	Meta PaginatedMeta             `json:"meta"`
	Data []ApplicationRegistration `json:"data"`
}

func (o *ListApplicationRegistrationsResponse) GetMeta() PaginatedMeta {
	if o == nil {
		return PaginatedMeta{}
	}
	return o.Meta
}

func (o *ListApplicationRegistrationsResponse) GetData() []ApplicationRegistration {
	if o == nil {
		return []ApplicationRegistration{}
	}
	return o.Data
}
