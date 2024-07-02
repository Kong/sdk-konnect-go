// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type ListControlPlanesResponse struct {
	// returns the pagination information
	Meta PaginatedMeta  `json:"meta"`
	Data []ControlPlane `json:"data"`
}

func (o *ListControlPlanesResponse) GetMeta() PaginatedMeta {
	if o == nil {
		return PaginatedMeta{}
	}
	return o.Meta
}

func (o *ListControlPlanesResponse) GetData() []ControlPlane {
	if o == nil {
		return []ControlPlane{}
	}
	return o.Data
}