# AIGatewayDatastoreUsageEntryType

The kind of entity holding the reference. A policy holds one in its `datastores` array. A
model holds one when its balancer is `semantic`. An MCP server, a model provider, or an
auth strategy holds one when it stores credentials in the Token Vault.

**`auth_strategy` requires a minimum runtime version of `2.3`**.

**`mcp_server` requires a minimum runtime version of `2.3`**.

**`model_provider` requires a minimum runtime version of `2.3`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayDatastoreUsageEntryTypePolicy

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayDatastoreUsageEntryType("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `AIGatewayDatastoreUsageEntryTypePolicy`        | policy                                          |
| `AIGatewayDatastoreUsageEntryTypeModel`         | model                                           |
| `AIGatewayDatastoreUsageEntryTypeMcpServer`     | mcp_server                                      |
| `AIGatewayDatastoreUsageEntryTypeModelProvider` | model_provider                                  |
| `AIGatewayDatastoreUsageEntryTypeAuthStrategy`  | auth_strategy                                   |