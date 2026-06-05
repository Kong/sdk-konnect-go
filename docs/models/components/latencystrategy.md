# LatencyStrategy

What metrics to use for latency. Available values are: `tpot` (time-per-output-token) and `e2e`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.LatencyStrategyE2e

// Open enum: custom values can be created with a direct type cast
custom := components.LatencyStrategy("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `LatencyStrategyE2e`  | e2e                   |
| `LatencyStrategyTpot` | tpot                  |