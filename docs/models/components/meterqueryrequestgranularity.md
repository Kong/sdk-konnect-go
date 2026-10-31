# MeterQueryRequestGranularity

The size of the time buckets to group the usage into. If not specified, the
usage is aggregated over the entire period.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MeterQueryRequestGranularityPt1M

// Open enum: custom values can be created with a direct type cast
custom := components.MeterQueryRequestGranularity("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `MeterQueryRequestGranularityPt1M` | PT1M                               |
| `MeterQueryRequestGranularityPt1H` | PT1H                               |
| `MeterQueryRequestGranularityP1D`  | P1D                                |
| `MeterQueryRequestGranularityP1M`  | P1M                                |