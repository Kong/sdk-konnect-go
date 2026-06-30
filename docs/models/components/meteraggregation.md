# MeterAggregation

The aggregation type to use for the meter.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MeterAggregationSum

// Open enum: custom values can be created with a direct type cast
custom := components.MeterAggregation("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `MeterAggregationSum`         | sum                           |
| `MeterAggregationCount`       | count                         |
| `MeterAggregationUniqueCount` | unique_count                  |
| `MeterAggregationAvg`         | avg                           |
| `MeterAggregationMin`         | min                           |
| `MeterAggregationMax`         | max                           |
| `MeterAggregationLatest`      | latest                        |