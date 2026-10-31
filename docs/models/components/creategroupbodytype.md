# CreateGroupBodyType

Type of the group. Immutable after creation.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateGroupBodyTypeGroup

// Open enum: custom values can be created with a direct type cast
custom := components.CreateGroupBodyType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `CreateGroupBodyTypeGroup` | group                      |
| `CreateGroupBodyTypeTeam`  | team                       |