# SkillStatus

The current validation status of the skill. `pending` means the content has not yet been read and validated.

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
| `SkillStatusValid`   | valid                |
| `SkillStatusInvalid` | invalid              |