# AICostUsageRequestMode

How the invocation was served: `stream`, `oneshot`, or `realtime`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostUsageRequestModeStream

// Open enum: custom values can be created with a direct type cast
custom := components.AICostUsageRequestMode("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `AICostUsageRequestModeStream`   | stream                           |
| `AICostUsageRequestModeOneshot`  | oneshot                          |
| `AICostUsageRequestModeRealtime` | realtime                         |