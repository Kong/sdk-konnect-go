# SessionBind

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SessionBindIP

// Open enum: custom values can be created with a direct type cast
custom := components.SessionBind("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `SessionBindIP`        | ip                     |
| `SessionBindScheme`    | scheme                 |
| `SessionBindUserAgent` | user-agent             |