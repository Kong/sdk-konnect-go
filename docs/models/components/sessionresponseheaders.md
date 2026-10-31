# SessionResponseHeaders

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SessionResponseHeadersAbsoluteTimeout

// Open enum: custom values can be created with a direct type cast
custom := components.SessionResponseHeaders("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `SessionResponseHeadersAbsoluteTimeout` | absolute-timeout                        |
| `SessionResponseHeadersAudience`        | audience                                |
| `SessionResponseHeadersID`              | id                                      |
| `SessionResponseHeadersIdlingTimeout`   | idling-timeout                          |
| `SessionResponseHeadersRollingTimeout`  | rolling-timeout                         |
| `SessionResponseHeadersSubject`         | subject                                 |
| `SessionResponseHeadersTimeout`         | timeout                                 |