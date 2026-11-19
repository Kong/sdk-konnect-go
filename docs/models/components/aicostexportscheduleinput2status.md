# AICostExportScheduleInput2Status

Whether the schedule creates exports.

Defaults to `active`, so a replacement that omits `status` resumes a paused schedule.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportScheduleInput2StatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportScheduleInput2Status("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `AICostExportScheduleInput2StatusActive` | active                                   |
| `AICostExportScheduleInput2StatusPaused` | paused                                   |