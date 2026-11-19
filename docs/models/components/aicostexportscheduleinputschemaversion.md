# AICostExportScheduleInputSchemaVersion

The schema version of each export's data files within the dataset, which fixes their columns and values. Each export that the schedule creates copies it, so a newer schema version doesn't change the schedule's exports.

Defaults to `1`, also once newer schema versions exist, so a replacement that omits `schema_version` sets `1`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportScheduleInputSchemaVersionOne
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `AICostExportScheduleInputSchemaVersionOne` | 1                                           |