// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetIdpTeamMappingsRequest struct {
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
}

func (o *GetIdpTeamMappingsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *GetIdpTeamMappingsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

type GetIdpTeamMappingsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A paginated list response for a collection of team mappings.
	TeamMappingResponse *components.TeamMappingResponse
}

func (o *GetIdpTeamMappingsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetIdpTeamMappingsResponse) GetTeamMappingResponse() *components.TeamMappingResponse {
	if o == nil {
		return nil
	}
	return o.TeamMappingResponse
}
