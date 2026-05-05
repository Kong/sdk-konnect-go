# VirtualClusterTopicAliasConflict

How to handle conflicts where an alias shadows a physical topic.
* warn - activate the alias but log a warning and set the conflict metric to 1.
* ignore - activate the alias silently.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.VirtualClusterTopicAliasConflictWarn

// Open enum: custom values can be created with a direct type cast
custom := components.VirtualClusterTopicAliasConflict("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `VirtualClusterTopicAliasConflictWarn`   | warn                                     |
| `VirtualClusterTopicAliasConflictIgnore` | ignore                                   |