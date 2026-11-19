# SkillSyncStatusStatus

The state of the most recent synchronization attempt.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SkillSyncStatusStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.SkillSyncStatusStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `SkillSyncStatusStatusPending` | pending                        |
| `SkillSyncStatusStatusSuccess` | success                        |
| `SkillSyncStatusStatusFailed`  | failed                         |