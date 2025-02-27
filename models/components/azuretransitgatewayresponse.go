// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"time"
)

type AzureTransitGatewayResponse struct {
	// Human-readable name of the transit gateway.
	Name string `json:"name"`
	// List of mappings from remote DNS server IP address sets to proxied internal domains, for a transit gateway
	// attachment.
	//
	DNSConfig                      []TransitGatewayDNSConfig        `json:"dns_config"`
	TransitGatewayAttachmentConfig AzureVNETPeeringAttachmentConfig `json:"transit_gateway_attachment_config"`
	ID                             string                           `json:"id"`
	// State of the transit gateway.
	State TransitGatewayState `json:"state"`
	// Monotonically-increasing version count of the transit gateway, to indicate the order of updates to the
	// transit gateway.
	//
	EntityVersion int64 `json:"entity_version"`
	// An RFC-3339 timestamp representation of transit gateway creation date.
	CreatedAt time.Time `json:"created_at"`
	// An RFC-3339 timestamp representation of transit gateway update date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (a AzureTransitGatewayResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AzureTransitGatewayResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AzureTransitGatewayResponse) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AzureTransitGatewayResponse) GetDNSConfig() []TransitGatewayDNSConfig {
	if o == nil {
		return []TransitGatewayDNSConfig{}
	}
	return o.DNSConfig
}

func (o *AzureTransitGatewayResponse) GetTransitGatewayAttachmentConfig() AzureVNETPeeringAttachmentConfig {
	if o == nil {
		return AzureVNETPeeringAttachmentConfig{}
	}
	return o.TransitGatewayAttachmentConfig
}

func (o *AzureTransitGatewayResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AzureTransitGatewayResponse) GetState() TransitGatewayState {
	if o == nil {
		return TransitGatewayState("")
	}
	return o.State
}

func (o *AzureTransitGatewayResponse) GetEntityVersion() int64 {
	if o == nil {
		return 0
	}
	return o.EntityVersion
}

func (o *AzureTransitGatewayResponse) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *AzureTransitGatewayResponse) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
