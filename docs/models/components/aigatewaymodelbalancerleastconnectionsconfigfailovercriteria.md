# AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteria

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

value := components.AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaError

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteria("custom_value")
```


## Values

| Name                                                                        | Value                                                                       |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaError`         | error                                                                       |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp400`       | http_400                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp401`       | http_401                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp402`       | http_402                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp403`       | http_403                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp404`       | http_404                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp405`       | http_405                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp408`       | http_408                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp409`       | http_409                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp413`       | http_413                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp415`       | http_415                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp422`       | http_422                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp429`       | http_429                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp500`       | http_500                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp502`       | http_502                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp503`       | http_503                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp504`       | http_504                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp522`       | http_522                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp529`       | http_529                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp5xx`       | http_5xx                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaInvalidHeader` | invalid_header                                                              |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaNonIdempotent` | non_idempotent                                                              |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaTimeout`       | timeout                                                                     |