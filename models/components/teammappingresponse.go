// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

// TeamMappingResponsePage - The page object.
type TeamMappingResponsePage struct {
	// Page number.
	Number *int64 `json:"number,omitempty"`
	// Page size.
	Size *int64 `json:"size,omitempty"`
	// Total number of results.
	Total *int64 `json:"total,omitempty"`
}

func (o *TeamMappingResponsePage) GetNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *TeamMappingResponsePage) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *TeamMappingResponsePage) GetTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.Total
}

// Meta - Contains pagination data.
type Meta struct {
	// The page object.
	Page *TeamMappingResponsePage `json:"page,omitempty"`
}

func (o *Meta) GetPage() *TeamMappingResponsePage {
	if o == nil {
		return nil
	}
	return o.Page
}

type TeamMappingResponseData struct {
	// Group names.
	Group *string `json:"group,omitempty"`
	// Team ID's that belong to the specified group.
	TeamIds []string `json:"team_ids,omitempty"`
}

func (o *TeamMappingResponseData) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *TeamMappingResponseData) GetTeamIds() []string {
	if o == nil {
		return nil
	}
	return o.TeamIds
}

type TeamMappingResponse struct {
	// Contains pagination data.
	Meta *Meta                     `json:"meta,omitempty"`
	Data []TeamMappingResponseData `json:"data,omitempty"`
}

func (o *TeamMappingResponse) GetMeta() *Meta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *TeamMappingResponse) GetData() []TeamMappingResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}