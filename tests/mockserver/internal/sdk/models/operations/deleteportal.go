// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/utils"
)

// DeletePortalForce - If true, the portal will be deleted, automatically deleting all API publications. If the force param is not set, the deletion will only succeed if there are no APIs currently published.
type DeletePortalForce string

const (
	DeletePortalForceTrue  DeletePortalForce = "true"
	DeletePortalForceFalse DeletePortalForce = "false"
)

func (e DeletePortalForce) ToPointer() *DeletePortalForce {
	return &e
}
func (e *DeletePortalForce) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "true":
		fallthrough
	case "false":
		*e = DeletePortalForce(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeletePortalForce: %v", v)
	}
}

type DeletePortalRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
	// If true, the portal will be deleted, automatically deleting all API publications. If the force param is not set, the deletion will only succeed if there are no APIs currently published.
	Force *DeletePortalForce `default:"false" queryParam:"style=form,explode=true,name=force"`
}

func (d DeletePortalRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DeletePortalRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DeletePortalRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *DeletePortalRequest) GetForce() *DeletePortalForce {
	if o == nil {
		return nil
	}
	return o.Force
}

type DeletePortalResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeletePortalResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
