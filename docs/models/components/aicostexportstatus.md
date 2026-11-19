# AICostExportStatus

Where the export is in its lifecycle. A client reads the export again with `GET /v1/ai-cost-management/exports/{exportId}` to follow its progress.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `AICostExportStatusPending`   | pending                       |
| `AICostExportStatusRunning`   | running                       |
| `AICostExportStatusSucceeded` | succeeded                     |
| `AICostExportStatusFailed`    | failed                        |
| `AICostExportStatusExpired`   | expired                       |