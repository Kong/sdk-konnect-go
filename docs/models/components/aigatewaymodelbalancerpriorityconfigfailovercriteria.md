# AIGatewayModelBalancerPriorityConfigFailoverCriteria

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayModelBalancerPriorityConfigFailoverCriteriaError

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelBalancerPriorityConfigFailoverCriteria("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaError`         | error                                                               |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp403`       | http_403                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp404`       | http_404                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp429`       | http_429                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp500`       | http_500                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp502`       | http_502                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp503`       | http_503                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp504`       | http_504                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaInvalidHeader` | invalid_header                                                      |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaNonIdempotent` | non_idempotent                                                      |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaTimeout`       | timeout                                                             |