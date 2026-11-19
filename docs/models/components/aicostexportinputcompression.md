# AICostExportInputCompression

How the data files are compressed.

`none` or `gzip` for `csv`, and `zstd` for `parquet`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportInputCompressionNone

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportInputCompression("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `AICostExportInputCompressionNone` | none                               |
| `AICostExportInputCompressionGzip` | gzip                               |
| `AICostExportInputCompressionZstd` | zstd                               |