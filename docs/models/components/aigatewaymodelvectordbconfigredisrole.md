# AIGatewayModelVectorDBConfigRedisRole

Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayModelVectorDBConfigRedisRoleAny

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayModelVectorDBConfigRedisRole("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `AIGatewayModelVectorDBConfigRedisRoleAny`    | any                                           |
| `AIGatewayModelVectorDBConfigRedisRoleMaster` | master                                        |
| `AIGatewayModelVectorDBConfigRedisRoleSlave`  | slave                                         |