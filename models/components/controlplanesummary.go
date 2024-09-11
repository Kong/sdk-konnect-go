// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"time"
)

// ControlPlaneSummaryClusterType - The ClusterType value of the cluster associated with the Control Plane.
type ControlPlaneSummaryClusterType string

const (
	ControlPlaneSummaryClusterTypeClusterTypeControlPlane         ControlPlaneSummaryClusterType = "CLUSTER_TYPE_CONTROL_PLANE"
	ControlPlaneSummaryClusterTypeClusterTypeHybrid               ControlPlaneSummaryClusterType = "CLUSTER_TYPE_HYBRID"
	ControlPlaneSummaryClusterTypeClusterTypeK8SIngressController ControlPlaneSummaryClusterType = "CLUSTER_TYPE_K8S_INGRESS_CONTROLLER"
	ControlPlaneSummaryClusterTypeClusterTypeControlPlaneGroup    ControlPlaneSummaryClusterType = "CLUSTER_TYPE_CONTROL_PLANE_GROUP"
	ControlPlaneSummaryClusterTypeClusterTypeServerless           ControlPlaneSummaryClusterType = "CLUSTER_TYPE_SERVERLESS"
)

func (e ControlPlaneSummaryClusterType) ToPointer() *ControlPlaneSummaryClusterType {
	return &e
}
func (e *ControlPlaneSummaryClusterType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CLUSTER_TYPE_CONTROL_PLANE":
		fallthrough
	case "CLUSTER_TYPE_HYBRID":
		fallthrough
	case "CLUSTER_TYPE_K8S_INGRESS_CONTROLLER":
		fallthrough
	case "CLUSTER_TYPE_CONTROL_PLANE_GROUP":
		fallthrough
	case "CLUSTER_TYPE_SERVERLESS":
		*e = ControlPlaneSummaryClusterType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ControlPlaneSummaryClusterType: %v", v)
	}
}

// ControlPlaneSummaryAuthType - The auth type value of the cluster associated with the Runtime Group.
type ControlPlaneSummaryAuthType string

const (
	ControlPlaneSummaryAuthTypePinnedClientCerts ControlPlaneSummaryAuthType = "pinned_client_certs"
	ControlPlaneSummaryAuthTypePkiClientCerts    ControlPlaneSummaryAuthType = "pki_client_certs"
)

func (e ControlPlaneSummaryAuthType) ToPointer() *ControlPlaneSummaryAuthType {
	return &e
}
func (e *ControlPlaneSummaryAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pinned_client_certs":
		fallthrough
	case "pki_client_certs":
		*e = ControlPlaneSummaryAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ControlPlaneSummaryAuthType: %v", v)
	}
}

// ControlPlaneSummaryConfig - CP configuration object for related access endpoints.
type ControlPlaneSummaryConfig struct {
	// Control Plane Endpoint.
	ControlPlaneEndpoint string `json:"control_plane_endpoint"`
	// Telemetry Endpoint.
	TelemetryEndpoint string `json:"telemetry_endpoint"`
	// The ClusterType value of the cluster associated with the Control Plane.
	ClusterType ControlPlaneSummaryClusterType `json:"cluster_type"`
	// The auth type value of the cluster associated with the Runtime Group.
	AuthType ControlPlaneSummaryAuthType `json:"auth_type"`
	// Whether the Control Plane can be used for cloud-gateways.
	CloudGateway bool `json:"cloud_gateway"`
	// Array of proxy URLs associated with reaching the data-planes connected to a control-plane.
	ProxyUrls []ProxyURL `json:"proxy_urls,omitempty"`
}

func (o *ControlPlaneSummaryConfig) GetControlPlaneEndpoint() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneEndpoint
}

func (o *ControlPlaneSummaryConfig) GetTelemetryEndpoint() string {
	if o == nil {
		return ""
	}
	return o.TelemetryEndpoint
}

func (o *ControlPlaneSummaryConfig) GetClusterType() ControlPlaneSummaryClusterType {
	if o == nil {
		return ControlPlaneSummaryClusterType("")
	}
	return o.ClusterType
}

func (o *ControlPlaneSummaryConfig) GetAuthType() ControlPlaneSummaryAuthType {
	if o == nil {
		return ControlPlaneSummaryAuthType("")
	}
	return o.AuthType
}

func (o *ControlPlaneSummaryConfig) GetCloudGateway() bool {
	if o == nil {
		return false
	}
	return o.CloudGateway
}

func (o *ControlPlaneSummaryConfig) GetProxyUrls() []ProxyURL {
	if o == nil {
		return nil
	}
	return o.ProxyUrls
}

// ControlPlaneSummary - The control plane object contains information about a Kong control plane.
type ControlPlaneSummary struct {
	// The control plane ID.
	ID string `json:"id"`
	// The name of the control plane.
	Name string `json:"name"`
	// The description of the control plane in Konnect.
	Description *string `json:"description,omitempty"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	Labels map[string]string `json:"labels,omitempty"`
	// CP configuration object for related access endpoints.
	Config    ControlPlaneSummaryConfig `json:"config"`
	CreatedAt time.Time                 `json:"created_at"`
	UpdatedAt time.Time                 `json:"updated_at"`
}

func (c ControlPlaneSummary) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ControlPlaneSummary) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ControlPlaneSummary) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ControlPlaneSummary) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ControlPlaneSummary) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ControlPlaneSummary) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *ControlPlaneSummary) GetConfig() ControlPlaneSummaryConfig {
	if o == nil {
		return ControlPlaneSummaryConfig{}
	}
	return o.Config
}

func (o *ControlPlaneSummary) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ControlPlaneSummary) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
