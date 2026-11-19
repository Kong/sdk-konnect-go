# AICostExportScheduleInputStatus

Whether the schedule creates exports.

Defaults to `active`, so a replacement that omits `status` resumes a paused schedule.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportScheduleInputStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportScheduleInputStatus("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `AICostExportScheduleInputStatusActive` | active                                  |
| `AICostExportScheduleInputStatusPaused` | paused                                  |