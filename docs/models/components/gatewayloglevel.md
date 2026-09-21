# GatewayLogLevel

A Kong Gateway log level

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.GatewayLogLevelDebug

// Open enum: custom values can be created with a direct type cast
custom := components.GatewayLogLevel("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `GatewayLogLevelDebug`  | debug                   |
| `GatewayLogLevelInfo`   | info                    |
| `GatewayLogLevelNotice` | notice                  |
| `GatewayLogLevelWarn`   | warn                    |
| `GatewayLogLevelError`  | error                   |
| `GatewayLogLevelCrit`   | crit                    |
| `GatewayLogLevelAlert`  | alert                   |
| `GatewayLogLevelEmerg`  | emerg                   |