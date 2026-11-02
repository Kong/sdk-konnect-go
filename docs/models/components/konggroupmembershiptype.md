# KongGroupMembershipType

Classifies the kind of member referenced by a group membership entry.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.KongGroupMembershipTypePrincipal

// Open enum: custom values can be created with a direct type cast
custom := components.KongGroupMembershipType("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `KongGroupMembershipTypePrincipal` | principal                          |
| `KongGroupMembershipTypeGroup`     | group                              |