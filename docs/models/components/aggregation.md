# Aggregation

The aggregation type to use for the meter.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AggregationSum

// Open enum: custom values can be created with a direct type cast
custom := components.Aggregation("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `AggregationSum`         | sum                      |
| `AggregationCount`       | count                    |
| `AggregationUniqueCount` | unique_count             |
| `AggregationAvg`         | avg                      |
| `AggregationMin`         | min                      |
| `AggregationMax`         | max                      |
| `AggregationLatest`      | latest                   |