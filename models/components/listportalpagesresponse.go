// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ListPortalPagesResponse - A paginated list of custom pages in a portal.
type ListPortalPagesResponse struct {
	Data []PortalPageInfo `json:"data"`
}

func (o *ListPortalPagesResponse) GetData() []PortalPageInfo {
	if o == nil {
		return []PortalPageInfo{}
	}
	return o.Data
}
