# Scopes

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ScopesGlobal

// Open enum: custom values can be created with a direct type cast
custom := components.Scopes("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `ScopesGlobal`         | global                 |
| `ScopesModels`         | models                 |
| `ScopesMcpServers`     | mcp-servers            |
| `ScopesAgents`         | agents                 |
| `ScopesConsumers`      | consumers              |
| `ScopesConsumerGroups` | consumer-groups        |