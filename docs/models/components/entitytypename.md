# EntityTypeName

The type of entity.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.EntityTypeNameAddOns

// Open enum: custom values can be created with a direct type cast
custom := components.EntityTypeName("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `EntityTypeNameAddOns`                    | Add Ons                                   |
| `EntityTypeNameApIs`                      | APIs                                      |
| `EntityTypeNameAPIProducts`               | API Products                              |
| `EntityTypeNameApplicationAuthStrategies` | Application Auth Strategies               |
| `EntityTypeNameAuditLogs`                 | Audit Logs                                |
| `EntityTypeNameControlPlanes`             | Control Planes                            |
| `EntityTypeNameDashboards`                | Dashboards                                |
| `EntityTypeNameDcrProviders`              | DCR Providers                             |
| `EntityTypeNameMeshControlPlanes`         | Mesh Control Planes                       |
| `EntityTypeNameNetworks`                  | Networks                                  |
| `EntityTypeNamePortals`                   | Portals                                   |
| `EntityTypeNameReports`                   | Reports                                   |
| `EntityTypeNameServiceHub`                | Service Hub                               |
| `EntityTypeNameAuthServers`               | Auth Servers                              |