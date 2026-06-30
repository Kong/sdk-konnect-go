# PlatformMetrics

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PlatformMetricsControlPlaneCount

// Open enum: custom values can be created with a direct type cast
custom := components.PlatformMetrics("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `PlatformMetricsControlPlaneCount` | control_plane_count                |
| `PlatformMetricsServiceCount`      | service_count                      |
| `PlatformMetricsRouteCount`        | route_count                        |
| `PlatformMetricsConsumerCount`     | consumer_count                     |
| `PlatformMetricsPluginCount`       | plugin_count                       |
| `PlatformMetricsNodeCount`         | node_count                         |