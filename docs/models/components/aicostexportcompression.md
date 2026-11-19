# AICostExportCompression

How the export's data files are compressed. Together with `format`, it sets the files' extension and content type.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportCompressionNone

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportCompression("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `AICostExportCompressionNone` | none                          |
| `AICostExportCompressionGzip` | gzip                          |
| `AICostExportCompressionZstd` | zstd                          |