// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/Kong/sdk-konnect-go/internal/utils"
)

type TargetWithoutParentsUpstream struct {
	ID *string `json:"id,omitempty"`
}

func (o *TargetWithoutParentsUpstream) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// TargetWithoutParents - A target is an ip address/hostname with a port that identifies an instance of a backend service. Every upstream can have many targets, and the targets can be dynamically added, modified, or deleted. Changes take effect on the fly. To disable a target, post a new one with `weight=0`; alternatively, use the `DELETE` convenience method to accomplish the same. The current target object definition is the one with the latest `created_at`.
type TargetWithoutParents struct {
	// Unix epoch when the resource was created.
	CreatedAt *float64 `json:"created_at,omitempty"`
	ID        *string  `json:"id,omitempty"`
	// An optional set of strings associated with the Target for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// The target address (ip or hostname) and port. If the hostname resolves to an SRV record, the `port` value will be overridden by the value from the DNS record.
	Target *string `json:"target,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *float64                      `json:"updated_at,omitempty"`
	Upstream  *TargetWithoutParentsUpstream `json:"upstream,omitempty"`
	// The weight this target gets within the upstream loadbalancer (`0`-`65535`). If the hostname resolves to an SRV record, the `weight` value will be overridden by the value from the DNS record.
	Weight *int64 `default:"100" json:"weight"`
}

func (t TargetWithoutParents) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TargetWithoutParents) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TargetWithoutParents) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TargetWithoutParents) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TargetWithoutParents) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *TargetWithoutParents) GetTarget() *string {
	if o == nil {
		return nil
	}
	return o.Target
}

func (o *TargetWithoutParents) GetUpdatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *TargetWithoutParents) GetUpstream() *TargetWithoutParentsUpstream {
	if o == nil {
		return nil
	}
	return o.Upstream
}

func (o *TargetWithoutParents) GetWeight() *int64 {
	if o == nil {
		return nil
	}
	return o.Weight
}
