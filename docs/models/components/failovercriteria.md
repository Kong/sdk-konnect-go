# FailoverCriteria

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
| `FailoverCriteriaHttp403`       | http_403                        |
| `FailoverCriteriaHttp404`       | http_404                        |
| `FailoverCriteriaHttp429`       | http_429                        |
| `FailoverCriteriaHttp500`       | http_500                        |
| `FailoverCriteriaHttp502`       | http_502                        |
| `FailoverCriteriaHttp503`       | http_503                        |
| `FailoverCriteriaHttp504`       | http_504                        |
| `FailoverCriteriaInvalidHeader` | invalid_header                  |
| `FailoverCriteriaNonIdempotent` | non_idempotent                  |
| `FailoverCriteriaTimeout`       | timeout                         |