// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"time"
)

// PortalSnippetResponse - Details about a page in a portal.
type PortalSnippetResponse struct {
	// Contains a unique identifier used for this resource.
	ID string `json:"id"`
	// The unique name of a snippet in a portal.
	Name string `json:"name"`
	// The display title of a snippet in a portal.
	Title *string `json:"title,omitempty"`
	// The renderable markdown content of a page in a portal.
	Content string `json:"content"`
	// Whether the resource is publicly accessible to non-authenticated users.
	Visibility VisibilityStatus `json:"visibility"`
	// Whether the resource is visible on a given portal. Defaults to unpublished.
	Status      PublishedStatus `json:"status"`
	Description *string         `json:"description,omitempty"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt time.Time `json:"created_at"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (p PortalSnippetResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PortalSnippetResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PortalSnippetResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PortalSnippetResponse) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PortalSnippetResponse) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *PortalSnippetResponse) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *PortalSnippetResponse) GetVisibility() VisibilityStatus {
	if o == nil {
		return VisibilityStatus("")
	}
	return o.Visibility
}

func (o *PortalSnippetResponse) GetStatus() PublishedStatus {
	if o == nil {
		return PublishedStatus("")
	}
	return o.Status
}

func (o *PortalSnippetResponse) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PortalSnippetResponse) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *PortalSnippetResponse) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
