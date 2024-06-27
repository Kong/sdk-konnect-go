// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/Kong/sdk-konnect-go/internal/utils"
)

type TargetUpstream struct {
	ID *string `json:"id,omitempty"`
}

func (o *TargetUpstream) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// Target - A target is an ip address/hostname with a port that identifies an instance of a backend service. Every upstream can have many targets, and the targets can be dynamically added, modified, or deleted. Changes take effect on the fly. To disable a target, post a new one with `weight=0`; alternatively, use the `DELETE` convenience method to accomplish the same. The current target object definition is the one with the latest `created_at`.
type Target struct {
	// An optional set of strings associated with the Target for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// The target address (ip or hostname) and port. If the hostname resolves to an SRV record, the `port` value will be overridden by the value from the DNS record.
	Target *string `json:"target,omitempty"`
	// The weight this target gets within the upstream loadbalancer (`0`-`65535`). If the hostname resolves to an SRV record, the `weight` value will be overridden by the value from the DNS record.
	Weight   *int64          `default:"100" json:"weight"`
	Upstream *TargetUpstream `json:"upstream,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *float64 `json:"created_at,omitempty"`
	ID        *string  `json:"id,omitempty"`
}

func (t Target) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *Target) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Target) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Target) GetTarget() *string {
	if o == nil {
		return nil
	}
	return o.Target
}

func (o *Target) GetWeight() *int64 {
	if o == nil {
		return nil
	}
	return o.Weight
}

func (o *Target) GetUpstream() *TargetUpstream {
	if o == nil {
		return nil
	}
	return o.Upstream
}

func (o *Target) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Target) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
