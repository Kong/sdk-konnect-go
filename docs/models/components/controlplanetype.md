# ControlPlaneType

The kind of control plane the linked consumer was found in.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ControlPlaneTypeAPIGateway

// Open enum: custom values can be created with a direct type cast
custom := components.ControlPlaneType("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `ControlPlaneTypeAPIGateway` | api_gateway                  |
| `ControlPlaneTypeAiGateway`  | ai_gateway                   |