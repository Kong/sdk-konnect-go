# PartialVectordbSentinelRole

Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PartialVectordbSentinelRoleAny

// Open enum: custom values can be created with a direct type cast
custom := components.PartialVectordbSentinelRole("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `PartialVectordbSentinelRoleAny`    | any                                 |
| `PartialVectordbSentinelRoleMaster` | master                              |
| `PartialVectordbSentinelRoleSlave`  | slave                               |