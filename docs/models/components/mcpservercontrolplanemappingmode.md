# MCPServerControlPlaneMappingMode

Whether the MCP Server deployment on this Control Plane is managed by Konnect defaults ('basic') or configured directly via Kong Operator custom resources ('advanced').

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MCPServerControlPlaneMappingModeBasic

// Open enum: custom values can be created with a direct type cast
custom := components.MCPServerControlPlaneMappingMode("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `MCPServerControlPlaneMappingModeBasic`    | basic                                      |
| `MCPServerControlPlaneMappingModeAdvanced` | advanced                                   |