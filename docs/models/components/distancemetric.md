# DistanceMetric

the distance metric to use for vector searches

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DistanceMetricCosine

// Open enum: custom values can be created with a direct type cast
custom := components.DistanceMetric("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `DistanceMetricCosine`    | cosine                    |
| `DistanceMetricEuclidean` | euclidean                 |