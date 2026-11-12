# PartialVectordbDistanceMetric

the distance metric to use for vector searches

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PartialVectordbDistanceMetricCosine

// Open enum: custom values can be created with a direct type cast
custom := components.PartialVectordbDistanceMetric("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `PartialVectordbDistanceMetricCosine`    | cosine                                   |
| `PartialVectordbDistanceMetricEuclidean` | euclidean                                |