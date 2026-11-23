# AICostExportScheduleInput2SchemaVersion

The schema version of each export's data files within the dataset, which fixes their columns and values. Each export that the schedule creates copies it, so a newer schema version doesn't change the schedule's exports.

Defaults to `1`, also once newer schema versions exist, so a replacement that omits `schema_version` sets `1`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportScheduleInput2SchemaVersionOne

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportScheduleInput2SchemaVersion(999)
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `AICostExportScheduleInput2SchemaVersionOne` | 1                                            |
| `AICostExportScheduleInput2SchemaVersionTwo` | 2                                            |