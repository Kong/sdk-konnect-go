# AICostExportFormat

The file format of the export's data files.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportFormatParquet

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportFormat("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `AICostExportFormatParquet` | parquet                     |
| `AICostExportFormatCsv`     | csv                         |