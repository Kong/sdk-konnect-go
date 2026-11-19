# AICostExport2Compression

How the export's data files are compressed. Together with `format`, it sets the files' extension and content type.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExport2CompressionNone

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExport2Compression("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `AICostExport2CompressionNone` | none                           |
| `AICostExport2CompressionGzip` | gzip                           |
| `AICostExport2CompressionZstd` | zstd                           |