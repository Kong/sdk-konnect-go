# AICostExportInputFormat

The file format of the data files.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportInputFormatParquet

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportInputFormat("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `AICostExportInputFormatParquet` | parquet                          |
| `AICostExportInputFormatCsv`     | csv                              |