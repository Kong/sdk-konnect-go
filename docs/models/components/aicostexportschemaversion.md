# AICostExportSchemaVersion

A version of an export's data files within its dataset: their columns and the values those columns hold. A change to either ships as a new schema version, so the files of a schema version never change. Each dataset counts its schema versions from `1` on its own. For `focus`, a schema version applies within the chosen FOCUS version. Each dataset offers these schema versions:

- `usage_events`: `1`.
- `focus`: `1`.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportSchemaVersionOne
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `AICostExportSchemaVersionOne` | 1                              |