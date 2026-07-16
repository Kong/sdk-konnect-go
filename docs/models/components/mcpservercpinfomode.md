# MCPServerCPInfoMode

The deployment mode of the MCP server on the associated Control Plane.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MCPServerCPInfoModeBasic

// Open enum: custom values can be created with a direct type cast
custom := components.MCPServerCPInfoMode("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `MCPServerCPInfoModeBasic`    | basic                         |
| `MCPServerCPInfoModeAdvanced` | advanced                      |