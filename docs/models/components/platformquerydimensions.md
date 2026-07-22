# PlatformQueryDimensions

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PlatformQueryDimensionsTime

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformQueryDimensions("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `PlatformQueryDimensionsTime`                 | time                                          |
| `PlatformQueryDimensionsControlPlane`         | control_plane                                 |
| `PlatformQueryDimensionsGatewayService`       | gateway_service                               |
| `PlatformQueryDimensionsRealm`                | realm                                         |
| `PlatformQueryDimensionsRoute`                | route                                         |
| `PlatformQueryDimensionsConsumer`             | consumer                                      |
| `PlatformQueryDimensionsPlugin`               | plugin                                        |
| `PlatformQueryDimensionsPluginName`           | plugin_name                                   |
| `PlatformQueryDimensionsPluginScope`          | plugin_scope                                  |
| `PlatformQueryDimensionsDataPlaneNodeVersion` | data_plane_node_version                       |
| `PlatformQueryDimensionsEnv`                  | env                                           |
| `PlatformQueryDimensionsTeam`                 | team                                          |
| `PlatformQueryDimensionsRegion`               | region                                        |
| `PlatformQueryDimensionsHostname`             | hostname                                      |