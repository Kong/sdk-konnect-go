// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// RoleName - The desired role.
type RoleName string

const (
	RoleNameAdmin                     RoleName = "Admin"
	RoleNameAppearanceMaintainer      RoleName = "Appearance Maintainer"
	RoleNameApplicationRegistration   RoleName = "Application Registration"
	RoleNameCertificateAdmin          RoleName = "Certificate Admin"
	RoleNameCloudGatewayClusterAdmin  RoleName = "Cloud Gateway Cluster Admin"
	RoleNameCloudGatewayClusterViewer RoleName = "Cloud Gateway Cluster Viewer"
	RoleNameConsumerAdmin             RoleName = "Consumer Admin"
	RoleNameConnector                 RoleName = "Connector"
	RoleNameCreator                   RoleName = "Creator"
	RoleNameDeployer                  RoleName = "Deployer"
	RoleNameDiscoveryAdmin            RoleName = "Discovery Admin"
	RoleNameDiscoveryViewer           RoleName = "Discovery Viewer"
	RoleNameGatewayServiceAdmin       RoleName = "Gateway Service Admin"
	RoleNameIntegrationAdmin          RoleName = "Integration Admin"
	RoleNameIntegrationViewer         RoleName = "Integration Viewer"
	RoleNameKeyAdmin                  RoleName = "Key Admin"
	RoleNameMaintainer                RoleName = "Maintainer"
	RoleNameNetworkAdmin              RoleName = "Network Admin"
	RoleNameNetworkCreator            RoleName = "Network Creator"
	RoleNameNetworkViewer             RoleName = "Network Viewer"
	RoleNamePluginAdmin               RoleName = "Plugin Admin"
	RoleNamePluginsAdmin              RoleName = "Plugins Admin"
	RoleNameProductPublisher          RoleName = "Product Publisher"
	RoleNamePublisher                 RoleName = "Publisher"
	RoleNameRouteAdmin                RoleName = "Route Admin"
	RoleNameSniAdmin                  RoleName = "SNI Admin"
	RoleNameScorecardAdmin            RoleName = "Scorecard Admin"
	RoleNameScorecardViewer           RoleName = "Scorecard Viewer"
	RoleNameServiceAdmin              RoleName = "Service Admin"
	RoleNameServiceCreator            RoleName = "Service Creator"
	RoleNameServiceViewer             RoleName = "Service Viewer"
	RoleNameUpstreamAdmin             RoleName = "Upstream Admin"
	RoleNameVaultAdmin                RoleName = "Vault Admin"
	RoleNameViewer                    RoleName = "Viewer"
)

func (e RoleName) ToPointer() *RoleName {
	return &e
}
func (e *RoleName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Admin":
		fallthrough
	case "Appearance Maintainer":
		fallthrough
	case "Application Registration":
		fallthrough
	case "Certificate Admin":
		fallthrough
	case "Cloud Gateway Cluster Admin":
		fallthrough
	case "Cloud Gateway Cluster Viewer":
		fallthrough
	case "Consumer Admin":
		fallthrough
	case "Connector":
		fallthrough
	case "Creator":
		fallthrough
	case "Deployer":
		fallthrough
	case "Discovery Admin":
		fallthrough
	case "Discovery Viewer":
		fallthrough
	case "Gateway Service Admin":
		fallthrough
	case "Integration Admin":
		fallthrough
	case "Integration Viewer":
		fallthrough
	case "Key Admin":
		fallthrough
	case "Maintainer":
		fallthrough
	case "Network Admin":
		fallthrough
	case "Network Creator":
		fallthrough
	case "Network Viewer":
		fallthrough
	case "Plugin Admin":
		fallthrough
	case "Plugins Admin":
		fallthrough
	case "Product Publisher":
		fallthrough
	case "Publisher":
		fallthrough
	case "Route Admin":
		fallthrough
	case "SNI Admin":
		fallthrough
	case "Scorecard Admin":
		fallthrough
	case "Scorecard Viewer":
		fallthrough
	case "Service Admin":
		fallthrough
	case "Service Creator":
		fallthrough
	case "Service Viewer":
		fallthrough
	case "Upstream Admin":
		fallthrough
	case "Vault Admin":
		fallthrough
	case "Viewer":
		*e = RoleName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RoleName: %v", v)
	}
}

// EntityTypeName - The type of entity.
type EntityTypeName string

const (
	EntityTypeNameApIs                      EntityTypeName = "APIs"
	EntityTypeNameAPIProducts               EntityTypeName = "API Products"
	EntityTypeNameApplicationAuthStrategies EntityTypeName = "Application Auth Strategies"
	EntityTypeNameAuditLogs                 EntityTypeName = "Audit Logs"
	EntityTypeNameControlPlanes             EntityTypeName = "Control Planes"
	EntityTypeNameDcrProviders              EntityTypeName = "DCR Providers"
	EntityTypeNameIdentity                  EntityTypeName = "Identity"
	EntityTypeNameMeshControlPlanes         EntityTypeName = "Mesh Control Planes"
	EntityTypeNameNetworks                  EntityTypeName = "Networks"
	EntityTypeNamePortals                   EntityTypeName = "Portals"
	EntityTypeNameServiceHub                EntityTypeName = "Service Hub"
)

func (e EntityTypeName) ToPointer() *EntityTypeName {
	return &e
}
func (e *EntityTypeName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "APIs":
		fallthrough
	case "API Products":
		fallthrough
	case "Application Auth Strategies":
		fallthrough
	case "Audit Logs":
		fallthrough
	case "Control Planes":
		fallthrough
	case "DCR Providers":
		fallthrough
	case "Identity":
		fallthrough
	case "Mesh Control Planes":
		fallthrough
	case "Networks":
		fallthrough
	case "Portals":
		fallthrough
	case "Service Hub":
		*e = EntityTypeName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EntityTypeName: %v", v)
	}
}

// AssignRole - An assigned role is a role that has been assigned to a user or team.
type AssignRole struct {
	// The desired role.
	RoleName *RoleName `json:"role_name,omitempty"`
	// The ID of the entity.
	EntityID *string `json:"entity_id,omitempty"`
	// The type of entity.
	EntityTypeName *EntityTypeName `json:"entity_type_name,omitempty"`
	EntityRegion   *EntityRegion   `json:"entity_region,omitempty"`
}

func (o *AssignRole) GetRoleName() *RoleName {
	if o == nil {
		return nil
	}
	return o.RoleName
}

func (o *AssignRole) GetEntityID() *string {
	if o == nil {
		return nil
	}
	return o.EntityID
}

func (o *AssignRole) GetEntityTypeName() *EntityTypeName {
	if o == nil {
		return nil
	}
	return o.EntityTypeName
}

func (o *AssignRole) GetEntityRegion() *EntityRegion {
	if o == nil {
		return nil
	}
	return o.EntityRegion
}
