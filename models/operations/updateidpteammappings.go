// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Kong/sdk-konnect-go/models/components"
	"net/http"
)

var UpdateIdpTeamMappingsServerList = []string{
	"https://global.api.konghq.com/",
}

type UpdateIdpTeamMappingsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A paginated list response for a collection of team mappings.
	TeamMappingCollection *components.TeamMappingCollection
}

func (o *UpdateIdpTeamMappingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateIdpTeamMappingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateIdpTeamMappingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateIdpTeamMappingsResponse) GetTeamMappingCollection() *components.TeamMappingCollection {
	if o == nil {
		return nil
	}
	return o.TeamMappingCollection
}
