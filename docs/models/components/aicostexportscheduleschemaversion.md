# AICostExportScheduleSchemaVersion

The schema version of each export's data files within the dataset, which fixes their columns and values.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportScheduleSchemaVersionOne

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportScheduleSchemaVersion(999)
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `AICostExportScheduleSchemaVersionOne` | 1                                      |
| `AICostExportScheduleSchemaVersionTwo` | 2                                      |