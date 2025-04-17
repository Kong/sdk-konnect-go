// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DegraphqlRouteService struct {
	ID *string `json:"id,omitempty"`
}

func (o *DegraphqlRouteService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type DegraphqlRoute struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64                 `json:"created_at,omitempty"`
	ID        *string                `json:"id,omitempty"`
	Methods   []string               `json:"methods,omitempty"`
	Query     string                 `json:"query"`
	Service   *DegraphqlRouteService `json:"service"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	URI       string `json:"uri"`
}

func (o *DegraphqlRoute) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *DegraphqlRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DegraphqlRoute) GetMethods() []string {
	if o == nil {
		return nil
	}
	return o.Methods
}

func (o *DegraphqlRoute) GetQuery() string {
	if o == nil {
		return ""
	}
	return o.Query
}

func (o *DegraphqlRoute) GetService() *DegraphqlRouteService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *DegraphqlRoute) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *DegraphqlRoute) GetURI() string {
	if o == nil {
		return ""
	}
	return o.URI
}
