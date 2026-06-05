# AIGatewayModelBalancerLowestLatencyConfigFailoverCriteria

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaError

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelBalancerLowestLatencyConfigFailoverCriteria("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaError`         | error                                                                    |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp403`       | http_403                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp404`       | http_404                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp429`       | http_429                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp500`       | http_500                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp502`       | http_502                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp503`       | http_503                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp504`       | http_504                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaInvalidHeader` | invalid_header                                                           |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaNonIdempotent` | non_idempotent                                                           |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaTimeout`       | timeout                                                                  |