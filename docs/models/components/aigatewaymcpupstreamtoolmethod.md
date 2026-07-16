# AIGatewayMCPUpstreamToolMethod

When provided, the method of the exported API, which must match the route's methods.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayMCPUpstreamToolMethodDelete

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayMCPUpstreamToolMethod("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `AIGatewayMCPUpstreamToolMethodDelete` | DELETE                                 |
| `AIGatewayMCPUpstreamToolMethodGet`    | GET                                    |
| `AIGatewayMCPUpstreamToolMethodPatch`  | PATCH                                  |
| `AIGatewayMCPUpstreamToolMethodPost`   | POST                                   |
| `AIGatewayMCPUpstreamToolMethodPut`    | PUT                                    |