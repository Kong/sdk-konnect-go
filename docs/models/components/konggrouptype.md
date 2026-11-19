# KongGroupType

Classifies the kind of group. There are three group types: `group`,
`team`, and `project`. `group` is the default policy-attachment point
for access across principals. `team` is a node of the organization
chart: its members are principals of type `user` and other `team`
groups, and it may carry an `owner` custom identity referencing a
principal of type `user`. `project` groups `user` and `agent`
principals; projects cannot contain groups or be members of groups.
When omitted or null, the group is a `group`.

Teams come from the organization import; of the three group types,
only `group` and `project` are mutable through this API.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.KongGroupTypeGroup

// Open enum: custom values can be created with a direct type cast
custom := components.KongGroupType("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `KongGroupTypeGroup` | group                |
| `KongGroupTypeTeam`  | team                 |