// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/utils"
)

// DeleteConfigStoreForce - If true, delete specified config store and all secrets, even if there are secrets linked to the config store If false, do not allow deletion if there are secrets linked to the config store
type DeleteConfigStoreForce string

const (
	DeleteConfigStoreForceTrue  DeleteConfigStoreForce = "true"
	DeleteConfigStoreForceFalse DeleteConfigStoreForce = "false"
)

func (e DeleteConfigStoreForce) ToPointer() *DeleteConfigStoreForce {
	return &e
}
func (e *DeleteConfigStoreForce) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "true":
		fallthrough
	case "false":
		*e = DeleteConfigStoreForce(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteConfigStoreForce: %v", v)
	}
}

type DeleteConfigStoreRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Config Store identifier
	ConfigStoreID string `pathParam:"style=simple,explode=false,name=configStoreId"`
	// If true, delete specified config store and all secrets, even if there are secrets linked to the config store If false, do not allow deletion if there are secrets linked to the config store
	Force *DeleteConfigStoreForce `default:"false" queryParam:"style=form,explode=true,name=force"`
}

func (d DeleteConfigStoreRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DeleteConfigStoreRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DeleteConfigStoreRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *DeleteConfigStoreRequest) GetConfigStoreID() string {
	if o == nil {
		return ""
	}
	return o.ConfigStoreID
}

func (o *DeleteConfigStoreRequest) GetForce() *DeleteConfigStoreForce {
	if o == nil {
		return nil
	}
	return o.Force
}

type DeleteConfigStoreResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteConfigStoreResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
