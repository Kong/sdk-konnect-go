# AICostExportScheduleInputCadence

The length of each period, which sets the range of each export.

A replacement can't change it.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportScheduleInputCadenceDaily

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportScheduleInputCadence("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `AICostExportScheduleInputCadenceDaily`   | daily                                     |
| `AICostExportScheduleInputCadenceWeekly`  | weekly                                    |
| `AICostExportScheduleInputCadenceMonthly` | monthly                                   |