# MCPResourceRemoteMCPServerPayloadType

The type of the MCP resource source.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MCPResourceRemoteMCPServerPayloadTypeAPICatalog

// Open enum: custom values can be created with a direct type cast
custom := components.MCPResourceRemoteMCPServerPayloadType("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `MCPResourceRemoteMCPServerPayloadTypeAPICatalog`      | api_catalog                                            |
| `MCPResourceRemoteMCPServerPayloadTypeRaw`             | raw                                                    |
| `MCPResourceRemoteMCPServerPayloadTypeRemoteMcpServer` | remote_mcp_server                                      |