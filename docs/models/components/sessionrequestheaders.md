# SessionRequestHeaders

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SessionRequestHeadersAbsoluteTimeout

// Open enum: custom values can be created with a direct type cast
custom := components.SessionRequestHeaders("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `SessionRequestHeadersAbsoluteTimeout` | absolute-timeout                       |
| `SessionRequestHeadersAudience`        | audience                               |
| `SessionRequestHeadersID`              | id                                     |
| `SessionRequestHeadersIdlingTimeout`   | idling-timeout                         |
| `SessionRequestHeadersRollingTimeout`  | rolling-timeout                        |
| `SessionRequestHeadersSubject`         | subject                                |
| `SessionRequestHeadersTimeout`         | timeout                                |