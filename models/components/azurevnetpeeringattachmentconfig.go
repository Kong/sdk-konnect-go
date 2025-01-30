// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type AzureVNETPeeringAttachmentType string

const (
	AzureVNETPeeringAttachmentTypeAzureVnetPeeringAttachment AzureVNETPeeringAttachmentType = "azure-vnet-peering-attachment"
)

func (e AzureVNETPeeringAttachmentType) ToPointer() *AzureVNETPeeringAttachmentType {
	return &e
}
func (e *AzureVNETPeeringAttachmentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "azure-vnet-peering-attachment":
		*e = AzureVNETPeeringAttachmentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AzureVNETPeeringAttachmentType: %v", v)
	}
}

type AzureVNETPeeringAttachmentConfig struct {
	Kind AzureVNETPeeringAttachmentType `json:"kind"`
	// Tenant ID for the Azure VNET Peering attachment.
	TenantID string `json:"tenant_id"`
	// Subscription ID for the Azure VNET Peering attachment.
	SubscriptionID string `json:"subscription_id"`
	// Resource Group Name for the Azure VNET Peering attachment.
	ResourceGroupName string `json:"resource_group_name"`
	// VNET Name for the Azure VNET Peering attachment.
	VnetName string `json:"vnet_name"`
}

func (o *AzureVNETPeeringAttachmentConfig) GetKind() AzureVNETPeeringAttachmentType {
	if o == nil {
		return AzureVNETPeeringAttachmentType("")
	}
	return o.Kind
}

func (o *AzureVNETPeeringAttachmentConfig) GetTenantID() string {
	if o == nil {
		return ""
	}
	return o.TenantID
}

func (o *AzureVNETPeeringAttachmentConfig) GetSubscriptionID() string {
	if o == nil {
		return ""
	}
	return o.SubscriptionID
}

func (o *AzureVNETPeeringAttachmentConfig) GetResourceGroupName() string {
	if o == nil {
		return ""
	}
	return o.ResourceGroupName
}

func (o *AzureVNETPeeringAttachmentConfig) GetVnetName() string {
	if o == nil {
		return ""
	}
	return o.VnetName
}
