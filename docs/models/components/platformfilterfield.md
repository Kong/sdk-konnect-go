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
| `PlatformFilterFieldGatewayService`       | gateway_service                           |
| `PlatformFilterFieldPlugin`               | plugin                                    |
| `PlatformFilterFieldPluginName`           | plugin_name                               |
| `PlatformFilterFieldPluginScope`          | plugin_scope                              |
| `PlatformFilterFieldRealm`                | realm                                     |
| `PlatformFilterFieldRoute`                | route                                     |