# AIGatewayModelBalancerRoundRobinConfigFailoverCriteria

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaError

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelBalancerRoundRobinConfigFailoverCriteria("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaError`         | error                                                                 |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp403`       | http_403                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp404`       | http_404                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp429`       | http_429                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp500`       | http_500                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp502`       | http_502                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp503`       | http_503                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp504`       | http_504                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaInvalidHeader` | invalid_header                                                        |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaNonIdempotent` | non_idempotent                                                        |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaTimeout`       | timeout                                                               |