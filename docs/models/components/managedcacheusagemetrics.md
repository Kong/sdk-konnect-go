# ManagedCacheUsageMetrics

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ManagedCacheUsageMetricsCacheEvictionRate

// Open enum: custom values can be created with a direct type cast
custom := components.ManagedCacheUsageMetrics("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `ManagedCacheUsageMetricsCacheEvictionRate`         | cache_eviction_rate                                 |
| `ManagedCacheUsageMetricsCacheExpirationRate`       | cache_expiration_rate                               |
| `ManagedCacheUsageMetricsCacheItemsAverage`         | cache_items_average                                 |
| `ManagedCacheUsageMetricsCacheMemoryUtilizationMax` | cache_memory_utilization_max                        |