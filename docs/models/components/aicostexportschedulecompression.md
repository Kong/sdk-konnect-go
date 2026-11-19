# AICostExportScheduleCompression

How each export's data files are compressed.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportScheduleCompressionNone

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportScheduleCompression("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `AICostExportScheduleCompressionNone` | none                                  |
| `AICostExportScheduleCompressionGzip` | gzip                                  |
| `AICostExportScheduleCompressionZstd` | zstd                                  |