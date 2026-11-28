# AIGatewayIdentityProviderOpenIDConnectResponseSentinelRole

Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectResponseSentinelRoleAny

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectResponseSentinelRole("custom_value")
```


## Values

| Name                                                               | Value                                                              |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `AIGatewayIdentityProviderOpenIDConnectResponseSentinelRoleAny`    | any                                                                |
| `AIGatewayIdentityProviderOpenIDConnectResponseSentinelRoleMaster` | master                                                             |
| `AIGatewayIdentityProviderOpenIDConnectResponseSentinelRoleSlave`  | slave                                                              |