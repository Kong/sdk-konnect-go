# AIGatewayModelBalancerLowestUsageConfigFailoverCriteria

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
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp403`       | http_403                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp404`       | http_404                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp429`       | http_429                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp500`       | http_500                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp502`       | http_502                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp503`       | http_503                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaHttp504`       | http_504                                                               |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaInvalidHeader` | invalid_header                                                         |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaNonIdempotent` | non_idempotent                                                         |
| `AIGatewayModelBalancerLowestUsageConfigFailoverCriteriaTimeout`       | timeout                                                                |