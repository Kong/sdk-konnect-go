// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

// GetTeamsTeamIDSystemAccountsFilter - Filter system accounts returned in the response.
type GetTeamsTeamIDSystemAccountsFilter struct {
	// Filters on the given string field value by either exact or fuzzy match.
	Name *components.StringFieldFilter `queryParam:"name=name"`
}

func (o *GetTeamsTeamIDSystemAccountsFilter) GetName() *components.StringFieldFilter {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetTeamsTeamIDSystemAccountsRequest struct {
	// ID of the team.
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Filter system accounts returned in the response.
	Filter *GetTeamsTeamIDSystemAccountsFilter `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (o *GetTeamsTeamIDSystemAccountsRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *GetTeamsTeamIDSystemAccountsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *GetTeamsTeamIDSystemAccountsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *GetTeamsTeamIDSystemAccountsRequest) GetFilter() *GetTeamsTeamIDSystemAccountsFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type GetTeamsTeamIDSystemAccountsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A paginated list response for a collection of system accounts.
	SystemAccountCollection *components.SystemAccountCollection
}

func (o *GetTeamsTeamIDSystemAccountsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTeamsTeamIDSystemAccountsResponse) GetSystemAccountCollection() *components.SystemAccountCollection {
	if o == nil {
		return nil
	}
	return o.SystemAccountCollection
}
