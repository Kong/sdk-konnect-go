// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"time"
)

// APIVersionResponseAPISpecType - The type of specification being stored. This allows us to render the specification correctly.
type APIVersionResponseAPISpecType string

const (
	APIVersionResponseAPISpecTypeOas2     APIVersionResponseAPISpecType = "oas2"
	APIVersionResponseAPISpecTypeOas3     APIVersionResponseAPISpecType = "oas3"
	APIVersionResponseAPISpecTypeAsyncapi APIVersionResponseAPISpecType = "asyncapi"
)

func (e APIVersionResponseAPISpecType) ToPointer() *APIVersionResponseAPISpecType {
	return &e
}
func (e *APIVersionResponseAPISpecType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oas2":
		fallthrough
	case "oas3":
		fallthrough
	case "asyncapi":
		*e = APIVersionResponseAPISpecType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APIVersionResponseAPISpecType: %v", v)
	}
}

type APIVersionResponseValidationMessages struct {
	Message string `json:"message"`
}

func (o *APIVersionResponseValidationMessages) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

type APIVersionResponseSpec struct {
	// The raw content of your API spec, in json or yaml format (OpenAPI or AsyncAPI).
	//
	Content *string                        `json:"content,omitempty"`
	Type    *APIVersionResponseAPISpecType `json:"type,omitempty"`
	// The errors that occurred while parsing the API version spec.
	ValidationMessages []APIVersionResponseValidationMessages `json:"validation_messages,omitempty"`
}

func (o *APIVersionResponseSpec) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *APIVersionResponseSpec) GetType() *APIVersionResponseAPISpecType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *APIVersionResponseSpec) GetValidationMessages() []APIVersionResponseValidationMessages {
	if o == nil {
		return nil
	}
	return o.ValidationMessages
}

// APIVersionResponse - API version (OpenAPI or AsyncAPI)
type APIVersionResponse struct {
	// The API version identifier.
	ID string `json:"id"`
	// The version of the api.
	Version string                  `json:"version"`
	Spec    *APIVersionResponseSpec `json:"spec,omitempty"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt time.Time `json:"created_at"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (a APIVersionResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *APIVersionResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *APIVersionResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *APIVersionResponse) GetVersion() string {
	if o == nil {
		return ""
	}
	return o.Version
}

func (o *APIVersionResponse) GetSpec() *APIVersionResponseSpec {
	if o == nil {
		return nil
	}
	return o.Spec
}

func (o *APIVersionResponse) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *APIVersionResponse) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
