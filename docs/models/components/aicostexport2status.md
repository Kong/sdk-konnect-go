# AICostExport2Status

Where the export is in its lifecycle. A client reads the export again with `GET /v1/ai-cost-management/exports/{exportId}` to follow its progress.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExport2StatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExport2Status("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `AICostExport2StatusPending`   | pending                        |
| `AICostExport2StatusRunning`   | running                        |
| `AICostExport2StatusSucceeded` | succeeded                      |
| `AICostExport2StatusFailed`    | failed                         |
| `AICostExport2StatusExpired`   | expired                        |