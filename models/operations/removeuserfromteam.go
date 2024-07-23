// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

var RemoveUserFromTeamServerList = []string{
	"https://global.api.konghq.com/",
}

type RemoveUserFromTeamRequest struct {
	// User ID
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
	// Team ID.
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
}

func (o *RemoveUserFromTeamRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *RemoveUserFromTeamRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

type RemoveUserFromTeamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RemoveUserFromTeamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveUserFromTeamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveUserFromTeamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
