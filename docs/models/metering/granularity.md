# Granularity

The size of the time buckets to group the usage into. If not specified, the
usage is aggregated over the entire period.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.GranularityPt1M

// Open enum: custom values can be created with a direct type cast
custom := metering.Granularity("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `GranularityPt1M` | PT1M              |
| `GranularityPt1H` | PT1H              |
| `GranularityP1D`  | P1D               |
| `GranularityP1M`  | P1M               |