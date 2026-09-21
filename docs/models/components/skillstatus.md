# SkillStatus

The current validation status of the skill.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SkillStatusValid

// Open enum: custom values can be created with a direct type cast
custom := components.SkillStatus("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `SkillStatusValid`   | valid                |
| `SkillStatusInvalid` | invalid              |