// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"mockserver/internal/sdk/utils"
	"time"
)

// APIImplementationResponse - An entity that implements an API
type APIImplementationResponse struct {
	// Contains a unique identifier used for this resource.
	ID string `json:"id"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt time.Time `json:"created_at"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt time.Time `json:"updated_at"`
	// A Gateway service that implements an API
	Service *APIImplementationService `json:"service,omitempty"`
}

func (a APIImplementationResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *APIImplementationResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *APIImplementationResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *APIImplementationResponse) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *APIImplementationResponse) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *APIImplementationResponse) GetService() *APIImplementationService {
	if o == nil {
		return nil
	}
	return o.Service
}
