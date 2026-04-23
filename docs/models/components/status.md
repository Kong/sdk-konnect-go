# Status

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

value := components.StatusDeploying

// Open enum: custom values can be created with a direct type cast
custom := components.Status("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `StatusDeploying` | deploying         |
| `StatusHealthy`   | healthy           |
| `StatusPending`   | pending           |
| `StatusUnhealthy` | unhealthy         |
| `StatusUpgrading` | upgrading         |