# AIGatewayModelBalancerLowestLatencyConfigFailoverCriteria

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

value := components.AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaError

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelBalancerLowestLatencyConfigFailoverCriteria("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaError`         | error                                                                    |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp400`       | http_400                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp401`       | http_401                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp402`       | http_402                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp403`       | http_403                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp404`       | http_404                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp405`       | http_405                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp408`       | http_408                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp409`       | http_409                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp413`       | http_413                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp415`       | http_415                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp422`       | http_422                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp429`       | http_429                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp500`       | http_500                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp502`       | http_502                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp503`       | http_503                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp504`       | http_504                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp522`       | http_522                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp529`       | http_529                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaHttp5xx`       | http_5xx                                                                 |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaInvalidHeader` | invalid_header                                                           |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaNonIdempotent` | non_idempotent                                                           |
| `AIGatewayModelBalancerLowestLatencyConfigFailoverCriteriaTimeout`       | timeout                                                                  |