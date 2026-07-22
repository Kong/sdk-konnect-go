# PlatformFilterField

The field to filter.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PlatformFilterFieldControlPlane

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformFilterField("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `PlatformFilterFieldControlPlane`         | control_plane                             |
| `PlatformFilterFieldDataPlaneNodeVersion` | data_plane_node_version                   |
| `PlatformFilterFieldEnv`                  | env                                       |
| `PlatformFilterFieldGatewayService`       | gateway_service                           |
| `PlatformFilterFieldHostname`             | hostname                                  |
| `PlatformFilterFieldPlugin`               | plugin                                    |
| `PlatformFilterFieldPluginName`           | plugin_name                               |
| `PlatformFilterFieldPluginScope`          | plugin_scope                              |
| `PlatformFilterFieldRealm`                | realm                                     |
| `PlatformFilterFieldRegion`               | region                                    |
| `PlatformFilterFieldRoute`                | route                                     |
| `PlatformFilterFieldTeam`                 | team                                      |