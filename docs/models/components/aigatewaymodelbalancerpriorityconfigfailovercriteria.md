# AIGatewayModelBalancerPriorityConfigFailoverCriteria

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

value := components.AIGatewayModelBalancerPriorityConfigFailoverCriteriaError

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelBalancerPriorityConfigFailoverCriteria("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaError`         | error                                                               |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp400`       | http_400                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp401`       | http_401                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp402`       | http_402                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp403`       | http_403                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp404`       | http_404                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp405`       | http_405                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp408`       | http_408                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp409`       | http_409                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp413`       | http_413                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp415`       | http_415                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp422`       | http_422                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp429`       | http_429                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp500`       | http_500                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp502`       | http_502                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp503`       | http_503                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp504`       | http_504                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp522`       | http_522                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp529`       | http_529                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaHttp5xx`       | http_5xx                                                            |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaInvalidHeader` | invalid_header                                                      |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaNonIdempotent` | non_idempotent                                                      |
| `AIGatewayModelBalancerPriorityConfigFailoverCriteriaTimeout`       | timeout                                                             |