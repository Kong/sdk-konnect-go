# MCPServerStatusResponseStatus

Aggregated deployment status of the MCP server.
- `deploying` — single version with desired replicas not yet fully ready.
- `healthy` — single version running with no failing pods.
- `pending` — no deployment status has been reported yet.
- `unhealthy` — one or more failing pods across any version.
- `upgrading` — multiple versions running with no failing pods.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MCPServerStatusResponseStatusDeploying

// Open enum: custom values can be created with a direct type cast
custom := components.MCPServerStatusResponseStatus("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `MCPServerStatusResponseStatusDeploying` | deploying                                |
| `MCPServerStatusResponseStatusHealthy`   | healthy                                  |
| `MCPServerStatusResponseStatusPending`   | pending                                  |
| `MCPServerStatusResponseStatusUnhealthy` | unhealthy                                |
| `MCPServerStatusResponseStatusUpgrading` | upgrading                                |