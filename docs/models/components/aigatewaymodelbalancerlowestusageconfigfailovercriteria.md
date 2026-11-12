# AIGatewayModelBalancerLowestUsageConfigFailoverCriteria

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

value := components.AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaError

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelBalancerLowestUsageConfigFailoverCriteria("custom_value")
```


## Values

| Name                                                                   | Value                                                                  |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaError`         | error                                                                  |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp400`       | http_400                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp401`       | http_401                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp402`       | http_402                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp403`       | http_403                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp404`       | http_404                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp405`       | http_405                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp408`       | http_408                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp409`       | http_409                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp413`       | http_413                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp415`       | http_415                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp422`       | http_422                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp429`       | http_429                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp500`       | http_500                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp502`       | http_502                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp503`       | http_503                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp504`       | http_504                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp522`       | http_522                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp529`       | http_529                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp5xx`       | http_5xx                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaInvalidHeader` | invalid_header                                                         |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaNonIdempotent` | non_idempotent                                                         |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaTimeout`       | timeout                                                                |