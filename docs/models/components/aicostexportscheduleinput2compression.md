# AICostExportScheduleInput2Compression

How each export's data files are compressed.

`none` or `gzip` for `csv`, and `zstd` for `parquet`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportScheduleInput2CompressionNone

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportScheduleInput2Compression("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `AICostExportScheduleInput2CompressionNone` | none                                        |
| `AICostExportScheduleInput2CompressionGzip` | gzip                                        |
| `AICostExportScheduleInput2CompressionZstd` | zstd                                        |