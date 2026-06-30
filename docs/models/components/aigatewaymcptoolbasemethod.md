# AIGatewayMCPToolBaseMethod

For conversion-only and conversion-listener modes, the method of the exported API, which must match the route's methods.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayMCPToolBaseMethodDelete

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayMCPToolBaseMethod("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `AIGatewayMCPToolBaseMethodDelete` | DELETE                             |
| `AIGatewayMCPToolBaseMethodGet`    | GET                                |
| `AIGatewayMCPToolBaseMethodPatch`  | PATCH                              |
| `AIGatewayMCPToolBaseMethodPost`   | POST                               |
| `AIGatewayMCPToolBaseMethodPut`    | PUT                                |