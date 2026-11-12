# SkillStatus

The state of the most recent synchronization attempt.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SkillStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.SkillStatus("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `SkillStatusPending` | pending              |
| `SkillStatusSuccess` | success              |
| `SkillStatusFailed`  | failed               |