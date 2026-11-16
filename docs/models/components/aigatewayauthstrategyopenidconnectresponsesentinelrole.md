# AIGatewayAuthStrategyOpenIDConnectResponseSentinelRole

Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectResponseSentinelRoleAny

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectResponseSentinelRole("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectResponseSentinelRoleAny`    | any                                                            |
| `AIGatewayAuthStrategyOpenIDConnectResponseSentinelRoleMaster` | master                                                         |
| `AIGatewayAuthStrategyOpenIDConnectResponseSentinelRoleSlave`  | slave                                                          |