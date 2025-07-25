// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CreatePortalSnippetRequest - Create a snippet in a portal.
type CreatePortalSnippetRequest struct {
	// The unique name of a snippet in a portal.
	Name string `json:"name"`
	// The display title of a snippet in a portal.
	Title *string `json:"title,omitempty"`
	// The renderable markdown content of a page in a portal.
	Content string `json:"content"`
	// Whether a snippet is publicly accessible to non-authenticated users.
	// If not provided, the default_page_visibility value of the portal will be used.
	//
	Visibility *SnippetVisibilityStatus `json:"visibility,omitempty"`
	// Whether the resource is visible on a given portal. Defaults to unpublished.
	Status      *PublishedStatus `json:"status,omitempty"`
	Description *string          `json:"description,omitempty"`
}

func (o *CreatePortalSnippetRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreatePortalSnippetRequest) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *CreatePortalSnippetRequest) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *CreatePortalSnippetRequest) GetVisibility() *SnippetVisibilityStatus {
	if o == nil {
		return nil
	}
	return o.Visibility
}

func (o *CreatePortalSnippetRequest) GetStatus() *PublishedStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *CreatePortalSnippetRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
