# AIGatewayModelBalancerRoundRobinConfigFailoverCriteria

**`http_400` requires a minimum runtime version of `2.3`**.

**`http_401` requires a minimum runtime version of `2.3`**.

**`http_402` requires a minimum runtime version of `2.3`**.

**`http_405` requires a minimum runtime version of `2.3`**.

**`http_408` requires a minimum runtime version of `2.3`**.

**`http_409` requires a minimum runtime version of `2.3`**.

**`http_413` requires a minimum runtime version of `2.3`**.

**`http_415` requires a minimum runtime version of `2.3`**.

**`http_422` requires a minimum runtime version of `2.3`**.

**`http_522` requires a minimum runtime version of `2.3`**.

**`http_529` requires a minimum runtime version of `2.3`**.

**`http_5xx` requires a minimum runtime version of `2.3`**.

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
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp400`       | http_400                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp401`       | http_401                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp402`       | http_402                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp403`       | http_403                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp404`       | http_404                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp405`       | http_405                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp408`       | http_408                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp409`       | http_409                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp413`       | http_413                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp415`       | http_415                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp422`       | http_422                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp429`       | http_429                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp500`       | http_500                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp502`       | http_502                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp503`       | http_503                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp504`       | http_504                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp522`       | http_522                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp529`       | http_529                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaHttp5xx`       | http_5xx                                                              |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaInvalidHeader` | invalid_header                                                        |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaNonIdempotent` | non_idempotent                                                        |
| `AIGatewayModelBalancerRoundRobinConfigFailoverCriteriaTimeout`       | timeout                                                               |