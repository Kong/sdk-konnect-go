// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetOrganizationsMeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Me Organization
	MeOrganization *components.MeOrganization
}

func (o *GetOrganizationsMeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetOrganizationsMeResponse) GetMeOrganization() *components.MeOrganization {
	if o == nil {
		return nil
	}
	return o.MeOrganization
}
