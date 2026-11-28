# AIGatewayFailoverCriteria

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

value := components.AIGatewayFailoverCriteriaError

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayFailoverCriteria("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `AIGatewayFailoverCriteriaError`         | error                                    |
| `AIGatewayFailoverCriteriaHttp400`       | http_400                                 |
| `AIGatewayFailoverCriteriaHttp401`       | http_401                                 |
| `AIGatewayFailoverCriteriaHttp402`       | http_402                                 |
| `AIGatewayFailoverCriteriaHttp403`       | http_403                                 |
| `AIGatewayFailoverCriteriaHttp404`       | http_404                                 |
| `AIGatewayFailoverCriteriaHttp405`       | http_405                                 |
| `AIGatewayFailoverCriteriaHttp408`       | http_408                                 |
| `AIGatewayFailoverCriteriaHttp409`       | http_409                                 |
| `AIGatewayFailoverCriteriaHttp413`       | http_413                                 |
| `AIGatewayFailoverCriteriaHttp415`       | http_415                                 |
| `AIGatewayFailoverCriteriaHttp422`       | http_422                                 |
| `AIGatewayFailoverCriteriaHttp429`       | http_429                                 |
| `AIGatewayFailoverCriteriaHttp500`       | http_500                                 |
| `AIGatewayFailoverCriteriaHttp502`       | http_502                                 |
| `AIGatewayFailoverCriteriaHttp503`       | http_503                                 |
| `AIGatewayFailoverCriteriaHttp504`       | http_504                                 |
| `AIGatewayFailoverCriteriaHttp522`       | http_522                                 |
| `AIGatewayFailoverCriteriaHttp529`       | http_529                                 |
| `AIGatewayFailoverCriteriaHttp5xx`       | http_5xx                                 |
| `AIGatewayFailoverCriteriaInvalidHeader` | invalid_header                           |
| `AIGatewayFailoverCriteriaNonIdempotent` | non_idempotent                           |
| `AIGatewayFailoverCriteriaTimeout`       | timeout                                  |