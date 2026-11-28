# AICostExportInput2SchemaVersion

The schema version of the data files within the dataset, which fixes their columns and values.

Defaults to `1`, also once newer schema versions exist.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportInput2SchemaVersionOne

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportInput2SchemaVersion(999)
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `AICostExportInput2SchemaVersionOne` | 1                                    |
| `AICostExportInput2SchemaVersionTwo` | 2                                    |