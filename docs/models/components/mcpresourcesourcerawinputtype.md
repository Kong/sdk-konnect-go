# MCPResourceSourceRawInputType

The type of the MCP resource source.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MCPResourceSourceRawInputTypeAPICatalog

// Open enum: custom values can be created with a direct type cast
custom := components.MCPResourceSourceRawInputType("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `MCPResourceSourceRawInputTypeAPICatalog`      | api_catalog                                    |
| `MCPResourceSourceRawInputTypeRaw`             | raw                                            |
| `MCPResourceSourceRawInputTypeRemoteMcpServer` | remote_mcp_server                              |