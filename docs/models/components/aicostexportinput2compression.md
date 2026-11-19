# AICostExportInput2Compression

How the data files are compressed.

`none` or `gzip` for `csv`, and `zstd` for `parquet`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportInput2CompressionNone

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportInput2Compression("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `AICostExportInput2CompressionNone` | none                                |
| `AICostExportInput2CompressionGzip` | gzip                                |
| `AICostExportInput2CompressionZstd` | zstd                                |