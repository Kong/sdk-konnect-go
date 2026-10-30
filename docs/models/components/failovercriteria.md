# FailoverCriteria

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

value := components.FailoverCriteriaError

// Open enum: custom values can be created with a direct type cast
custom := components.FailoverCriteria("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `FailoverCriteriaError`         | error                           |
| `FailoverCriteriaHttp400`       | http_400                        |
| `FailoverCriteriaHttp401`       | http_401                        |
| `FailoverCriteriaHttp402`       | http_402                        |
| `FailoverCriteriaHttp403`       | http_403                        |
| `FailoverCriteriaHttp404`       | http_404                        |
| `FailoverCriteriaHttp405`       | http_405                        |
| `FailoverCriteriaHttp408`       | http_408                        |
| `FailoverCriteriaHttp409`       | http_409                        |
| `FailoverCriteriaHttp413`       | http_413                        |
| `FailoverCriteriaHttp415`       | http_415                        |
| `FailoverCriteriaHttp422`       | http_422                        |
| `FailoverCriteriaHttp429`       | http_429                        |
| `FailoverCriteriaHttp500`       | http_500                        |
| `FailoverCriteriaHttp502`       | http_502                        |
| `FailoverCriteriaHttp503`       | http_503                        |
| `FailoverCriteriaHttp504`       | http_504                        |
| `FailoverCriteriaHttp522`       | http_522                        |
| `FailoverCriteriaHttp529`       | http_529                        |
| `FailoverCriteriaHttp5xx`       | http_5xx                        |
| `FailoverCriteriaInvalidHeader` | invalid_header                  |
| `FailoverCriteriaNonIdempotent` | non_idempotent                  |
| `FailoverCriteriaTimeout`       | timeout                         |