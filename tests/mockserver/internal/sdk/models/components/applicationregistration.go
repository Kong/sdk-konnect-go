// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"mockserver/internal/sdk/utils"
	"time"
)

// ApplicationRegistrationAPI - Details about the API the application is registered to.
type ApplicationRegistrationAPI struct {
	// Contains a unique identifier used for this resource.
	ID string `json:"id"`
	// The name of the API the application is registered to.
	Name string `json:"name"`
	// The version of the API the application is registered to.
	Version *string `json:"version"`
}

func (o *ApplicationRegistrationAPI) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ApplicationRegistrationAPI) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ApplicationRegistrationAPI) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

// ApplicationRegistrationApplication - Details about the application the registration is part of.
type ApplicationRegistrationApplication struct {
	// Contains a unique identifier used for this resource.
	ID string `json:"id"`
	// The name of the application the registration is part of.
	Name string `json:"name"`
}

func (o *ApplicationRegistrationApplication) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ApplicationRegistrationApplication) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// ApplicationRegistration - A application's registration for a specific version of an API.
type ApplicationRegistration struct {
	// Contains a unique identifier used for this resource.
	ID string `json:"id"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt time.Time `json:"created_at"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt time.Time `json:"updated_at"`
	// The status of an application registration request. Each registration is linked to a single API, and application credentials will not grant access to the API until the registration is approved.
	// Pending, revoked, and rejected registrations will not provide access to the API.
	Status ApplicationRegistrationStatus `json:"status"`
	// Details about the API the application is registered to.
	API ApplicationRegistrationAPI `json:"api"`
	// Details about the application the registration is part of.
	Application ApplicationRegistrationApplication `json:"application"`
}

func (a ApplicationRegistration) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ApplicationRegistration) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ApplicationRegistration) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ApplicationRegistration) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ApplicationRegistration) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *ApplicationRegistration) GetStatus() ApplicationRegistrationStatus {
	if o == nil {
		return ApplicationRegistrationStatus("")
	}
	return o.Status
}

func (o *ApplicationRegistration) GetAPI() ApplicationRegistrationAPI {
	if o == nil {
		return ApplicationRegistrationAPI{}
	}
	return o.API
}

func (o *ApplicationRegistration) GetApplication() ApplicationRegistrationApplication {
	if o == nil {
		return ApplicationRegistrationApplication{}
	}
	return o.Application
}
