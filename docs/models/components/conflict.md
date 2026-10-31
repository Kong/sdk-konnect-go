# Conflict

How to inform the user about conflicts where multiple backend topics would map to the same virtual topic name.
* warn - log in the Event Gateway logs. Additionally, it sets knep_namespace_topic_conflict to 1.
* ignore - do not do anything. It does not cause knep_namespace_topic_conflict metric to be set to 1.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ConflictWarn

// Open enum: custom values can be created with a direct type cast
custom := components.Conflict("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `ConflictWarn`   | warn             |
| `ConflictIgnore` | ignore           |