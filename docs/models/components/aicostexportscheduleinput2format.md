# AICostExportScheduleInput2Format

The file format of each export's data files.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportScheduleInput2FormatParquet

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportScheduleInput2Format("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `AICostExportScheduleInput2FormatParquet` | parquet                                   |
| `AICostExportScheduleInput2FormatCsv`     | csv                                       |