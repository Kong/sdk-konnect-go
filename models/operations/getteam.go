// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

type GetTeamRequest struct {
	// The team ID
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
}

func (o *GetTeamRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

type GetTeamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response including a single team.
	Team *components.Team
}

func (o *GetTeamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTeamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTeamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTeamResponse) GetTeam() *components.Team {
	if o == nil {
		return nil
	}
	return o.Team
}
