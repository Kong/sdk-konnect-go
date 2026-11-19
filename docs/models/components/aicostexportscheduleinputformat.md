# AICostExportScheduleInputFormat

The file format of each export's data files.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportScheduleInputFormatParquet

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportScheduleInputFormat("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `AICostExportScheduleInputFormatParquet` | parquet                                  |
| `AICostExportScheduleInputFormatCsv`     | csv                                      |