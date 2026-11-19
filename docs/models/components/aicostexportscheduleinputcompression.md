# AICostExportScheduleInputCompression

How each export's data files are compressed.

`none` or `gzip` for `csv`, and `zstd` for `parquet`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportScheduleInputCompressionNone

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportScheduleInputCompression("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `AICostExportScheduleInputCompressionNone` | none                                       |
| `AICostExportScheduleInputCompressionGzip` | gzip                                       |
| `AICostExportScheduleInputCompressionZstd` | zstd                                       |