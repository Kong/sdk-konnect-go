# DataPlaneNodeLogLevelOperationStatus

Aggregate operation status for the request.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DataPlaneNodeLogLevelOperationStatusAccepted

// Open enum: custom values can be created with a direct type cast
custom := components.DataPlaneNodeLogLevelOperationStatus("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `DataPlaneNodeLogLevelOperationStatusAccepted`   | accepted                                         |
| `DataPlaneNodeLogLevelOperationStatusInProgress` | in_progress                                      |
| `DataPlaneNodeLogLevelOperationStatusCompleted`  | completed                                        |
| `DataPlaneNodeLogLevelOperationStatusFailed`     | failed                                           |