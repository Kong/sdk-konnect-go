# ClusterCacheStrategy

The strategy to use for the cluster cache. If set, the plugin will share cache with nodes configured with the same strategy backend. Currentlly only introspection cache is shared.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ClusterCacheStrategyOff

// Open enum: custom values can be created with a direct type cast
custom := components.ClusterCacheStrategy("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ClusterCacheStrategyOff`   | off                         |
| `ClusterCacheStrategyRedis` | redis                       |