# RegistryReferenceAgentSourceType

The agent-level source set through create and `PATCH /agents/{agentId}`. `manual`: the card is pasted through the version endpoints, and no remote host is configured. `remote`: the card is fetched from a host (see `hostname`) and refreshable.

This reflects only that agent-level setting. It does not reflect a source created through `/agents/{agentId}/source`, so an agent with an `ai_gateway` source there still reads `manual`.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.RegistryReferenceAgentSourceTypeManual
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `RegistryReferenceAgentSourceTypeManual` | manual                                   |
| `RegistryReferenceAgentSourceTypeRemote` | remote                                   |