# AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteria

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
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp403`       | http_403                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp404`       | http_404                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp429`       | http_429                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp500`       | http_500                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp502`       | http_502                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp503`       | http_503                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaHttp504`       | http_504                                                                    |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaInvalidHeader` | invalid_header                                                              |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaNonIdempotent` | non_idempotent                                                              |
| `AIGatewayModelBalancerLeastConnectionsConfigFailoverCriteriaTimeout`       | timeout                                                                     |