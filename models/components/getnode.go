// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type GetNodeCompatibilityStatus struct {
	State *string `json:"state,omitempty"`
}

func (o *GetNodeCompatibilityStatus) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

type GetNodeItem struct {
	ID                  *string                     `json:"id,omitempty"`
	Version             *string                     `json:"version,omitempty"`
	Hostname            *string                     `json:"hostname,omitempty"`
	LastPing            *int64                      `json:"last_ping,omitempty"`
	Type                *string                     `json:"type,omitempty"`
	CreatedAt           *int64                      `json:"created_at,omitempty"`
	UpdatedAt           *int64                      `json:"updated_at,omitempty"`
	ConfigHash          *string                     `json:"config_hash,omitempty"`
	CompatibilityStatus *GetNodeCompatibilityStatus `json:"compatibility_status,omitempty"`
}

func (o *GetNodeItem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetNodeItem) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *GetNodeItem) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *GetNodeItem) GetLastPing() *int64 {
	if o == nil {
		return nil
	}
	return o.LastPing
}

func (o *GetNodeItem) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetNodeItem) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *GetNodeItem) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *GetNodeItem) GetConfigHash() *string {
	if o == nil {
		return nil
	}
	return o.ConfigHash
}

func (o *GetNodeItem) GetCompatibilityStatus() *GetNodeCompatibilityStatus {
	if o == nil {
		return nil
	}
	return o.CompatibilityStatus
}

// GetNode - Example response
type GetNode struct {
	Item *GetNodeItem `json:"item,omitempty"`
}

func (o *GetNode) GetItem() *GetNodeItem {
	if o == nil {
		return nil
	}
	return o.Item
}
