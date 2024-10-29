// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// CreateControlPlaneRequestClusterType - The ClusterType value of the cluster associated with the Control Plane.
type CreateControlPlaneRequestClusterType string

const (
	CreateControlPlaneRequestClusterTypeClusterTypeControlPlane         CreateControlPlaneRequestClusterType = "CLUSTER_TYPE_CONTROL_PLANE"
	CreateControlPlaneRequestClusterTypeClusterTypeK8SIngressController CreateControlPlaneRequestClusterType = "CLUSTER_TYPE_K8S_INGRESS_CONTROLLER"
	CreateControlPlaneRequestClusterTypeClusterTypeControlPlaneGroup    CreateControlPlaneRequestClusterType = "CLUSTER_TYPE_CONTROL_PLANE_GROUP"
	CreateControlPlaneRequestClusterTypeClusterTypeServerless           CreateControlPlaneRequestClusterType = "CLUSTER_TYPE_SERVERLESS"
)

func (e CreateControlPlaneRequestClusterType) ToPointer() *CreateControlPlaneRequestClusterType {
	return &e
}
func (e *CreateControlPlaneRequestClusterType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CLUSTER_TYPE_CONTROL_PLANE":
		fallthrough
	case "CLUSTER_TYPE_K8S_INGRESS_CONTROLLER":
		fallthrough
	case "CLUSTER_TYPE_CONTROL_PLANE_GROUP":
		fallthrough
	case "CLUSTER_TYPE_SERVERLESS":
		*e = CreateControlPlaneRequestClusterType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateControlPlaneRequestClusterType: %v", v)
	}
}

// AuthType - The auth type value of the cluster associated with the Runtime Group.
type AuthType string

const (
	AuthTypePinnedClientCerts AuthType = "pinned_client_certs"
	AuthTypePkiClientCerts    AuthType = "pki_client_certs"
)

func (e AuthType) ToPointer() *AuthType {
	return &e
}
func (e *AuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pinned_client_certs":
		fallthrough
	case "pki_client_certs":
		*e = AuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthType: %v", v)
	}
}

// CreateControlPlaneRequest - The request schema for the create control plane request.
type CreateControlPlaneRequest struct {
	// The name of the control plane.
	Name string `json:"name"`
	// The description of the control plane in Konnect.
	Description *string `json:"description,omitempty"`
	// The ClusterType value of the cluster associated with the Control Plane.
	ClusterType *CreateControlPlaneRequestClusterType `json:"cluster_type,omitempty"`
	// The auth type value of the cluster associated with the Runtime Group.
	AuthType *AuthType `json:"auth_type,omitempty"`
	// Whether this control-plane can be used for cloud-gateways.
	CloudGateway *bool `json:"cloud_gateway,omitempty"`
	// Array of proxy URLs associated with reaching the data-planes connected to a control-plane.
	ProxyUrls []ProxyURL `json:"proxy_urls,omitempty"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	Labels map[string]string `json:"labels,omitempty"`
}

func (o *CreateControlPlaneRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateControlPlaneRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateControlPlaneRequest) GetClusterType() *CreateControlPlaneRequestClusterType {
	if o == nil {
		return nil
	}
	return o.ClusterType
}

func (o *CreateControlPlaneRequest) GetAuthType() *AuthType {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *CreateControlPlaneRequest) GetCloudGateway() *bool {
	if o == nil {
		return nil
	}
	return o.CloudGateway
}

func (o *CreateControlPlaneRequest) GetProxyUrls() []ProxyURL {
	if o == nil {
		return nil
	}
	return o.ProxyUrls
}

func (o *CreateControlPlaneRequest) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}
