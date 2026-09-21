# AIGatewayMCPConversionToolMethod

The HTTP method used when forwarding the request to the upstream API.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayMCPConversionToolMethodDelete

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayMCPConversionToolMethod("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `AIGatewayMCPConversionToolMethodDelete` | DELETE                                   |
| `AIGatewayMCPConversionToolMethodGet`    | GET                                      |
| `AIGatewayMCPConversionToolMethodPatch`  | PATCH                                    |
| `AIGatewayMCPConversionToolMethodPost`   | POST                                     |
| `AIGatewayMCPConversionToolMethodPut`    | PUT                                      |