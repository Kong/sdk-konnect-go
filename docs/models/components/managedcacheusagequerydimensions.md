# ManagedCacheUsageQueryDimensions

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ManagedCacheUsageQueryDimensionsControlPlane

// Open enum: custom values can be created with a direct type cast
custom := components.ManagedCacheUsageQueryDimensions("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `ManagedCacheUsageQueryDimensionsControlPlane`   | control_plane                                    |
| `ManagedCacheUsageQueryDimensionsDataPlaneGroup` | data_plane_group                                 |
| `ManagedCacheUsageQueryDimensionsManagedCache`   | managed_cache                                    |
| `ManagedCacheUsageQueryDimensionsNetwork`        | network                                          |
| `ManagedCacheUsageQueryDimensionsProvider`       | provider                                         |
| `ManagedCacheUsageQueryDimensionsProviderRegion` | provider_region                                  |
| `ManagedCacheUsageQueryDimensionsTime`           | time                                             |