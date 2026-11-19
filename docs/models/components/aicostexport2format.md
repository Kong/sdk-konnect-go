# AICostExport2Format

The file format of the export's data files.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExport2FormatParquet

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExport2Format("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `AICostExport2FormatParquet` | parquet                      |
| `AICostExport2FormatCsv`     | csv                          |