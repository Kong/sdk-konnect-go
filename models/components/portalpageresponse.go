// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"time"
)

// PortalPageResponse - Details about a page in a portal.
type PortalPageResponse struct {
	// Contains a unique identifier used for this resource.
	ID string `json:"id"`
	// The slug of a page in a portal, used to compute its full URL path within the portal hierarchy.
	// When a page has a `parent_page_id`, its full path is built by joining the parent’s slug with its own.
	// For example, if a parent page has the slug `slug1` and this page’s slug is `slug2`, the resulting path will be `/slug1/slug2`.
	// This enables nested page structures like `/slug1/slug2/slug3`.
	//
	Slug string `json:"slug"`
	// The title of a page in a portal.
	Title string `json:"title"`
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
	// Pages may be rendered as a tree of files.
	//
	// Specify the `id` of another page as the `parent_page_id` to add some hierarchy to your pages.
	//
	ParentPageID *string `json:"parent_page_id"`
}

func (p PortalPageResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PortalPageResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PortalPageResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PortalPageResponse) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *PortalPageResponse) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *PortalPageResponse) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *PortalPageResponse) GetVisibility() VisibilityStatus {
	if o == nil {
		return VisibilityStatus("")
	}
	return o.Visibility
}

func (o *PortalPageResponse) GetStatus() PublishedStatus {
	if o == nil {
		return PublishedStatus("")
	}
	return o.Status
}

func (o *PortalPageResponse) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PortalPageResponse) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *PortalPageResponse) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *PortalPageResponse) GetParentPageID() *string {
	if o == nil {
		return nil
	}
	return o.ParentPageID
}
