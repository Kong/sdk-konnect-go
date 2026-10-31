# SentinelRole

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SentinelRoleMaster

// Open enum: custom values can be created with a direct type cast
custom := components.SentinelRole("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `SentinelRoleMaster` | master               |
| `SentinelRoleSlave`  | slave                |
| `SentinelRoleAny`    | any                  |