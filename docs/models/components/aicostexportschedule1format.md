# AICostExportSchedule1Format

The file format of each export's data files.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportSchedule1FormatParquet

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportSchedule1Format("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `AICostExportSchedule1FormatParquet` | parquet                              |
| `AICostExportSchedule1FormatCsv`     | csv                                  |