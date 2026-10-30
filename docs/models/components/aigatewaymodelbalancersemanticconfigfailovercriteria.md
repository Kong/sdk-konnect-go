# AIGatewayModelBalancerSemanticConfigFailoverCriteria

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

value := components.AIGatewayModelBalancerSemanticConfigFailoverCriteriaError

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelBalancerSemanticConfigFailoverCriteria("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaError`         | error                                                               |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp400`       | http_400                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp401`       | http_401                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp402`       | http_402                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp403`       | http_403                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp404`       | http_404                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp405`       | http_405                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp408`       | http_408                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp409`       | http_409                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp413`       | http_413                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp415`       | http_415                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp422`       | http_422                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp429`       | http_429                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp500`       | http_500                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp502`       | http_502                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp503`       | http_503                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp504`       | http_504                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp522`       | http_522                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp529`       | http_529                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp5xx`       | http_5xx                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaInvalidHeader` | invalid_header                                                      |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaNonIdempotent` | non_idempotent                                                      |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaTimeout`       | timeout                                                             |