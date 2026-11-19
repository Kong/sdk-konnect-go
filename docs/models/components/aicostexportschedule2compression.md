# AICostExportSchedule2Compression

How each export's data files are compressed.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportSchedule2CompressionNone

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportSchedule2Compression("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `AICostExportSchedule2CompressionNone` | none                                   |
| `AICostExportSchedule2CompressionGzip` | gzip                                   |
| `AICostExportSchedule2CompressionZstd` | zstd                                   |