# AIGatewayMCPServerUpstreamServerServerConfigStrategy

The strategy for the session. If the value is 'client', the session is encrypted into MCP session id assigned to the client. If the value is not 'client', the session is stored in the configured database.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayMCPServerUpstreamServerServerConfigStrategyClient

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayMCPServerUpstreamServerServerConfigStrategy("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `AIGatewayMCPServerUpstreamServerServerConfigStrategyClient` | client                                                       |
| `AIGatewayMCPServerUpstreamServerServerConfigStrategyRedis`  | redis                                                        |