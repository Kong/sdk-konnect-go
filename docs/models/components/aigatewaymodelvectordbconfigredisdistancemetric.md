# AIGatewayModelVectorDBConfigRedisDistanceMetric

the distance metric to use for vector searches

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayModelVectorDBConfigRedisDistanceMetricCosine

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelVectorDBConfigRedisDistanceMetric("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `AIGatewayModelVectorDBConfigRedisDistanceMetricCosine`    | cosine                                                     |
| `AIGatewayModelVectorDBConfigRedisDistanceMetricEuclidean` | euclidean                                                  |