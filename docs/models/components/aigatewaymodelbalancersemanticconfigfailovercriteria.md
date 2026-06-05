# AIGatewayModelBalancerSemanticConfigFailoverCriteria

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
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp403`       | http_403                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp404`       | http_404                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp429`       | http_429                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp500`       | http_500                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp502`       | http_502                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp503`       | http_503                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaHttp504`       | http_504                                                            |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaInvalidHeader` | invalid_header                                                      |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaNonIdempotent` | non_idempotent                                                      |
| `AIGatewayModelBalancerSemanticConfigFailoverCriteriaTimeout`       | timeout                                                             |