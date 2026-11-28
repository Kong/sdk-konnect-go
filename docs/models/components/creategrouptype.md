# CreateGroupType

Classifies the kind of group to create. There are three group types
(`group`, `team`, `project`), but only `group` and `project` are
mutable through this API; teams come from the organization import.
When omitted or null, the group is a `group`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateGroupTypeGroup
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `CreateGroupTypeGroup` | group                  |