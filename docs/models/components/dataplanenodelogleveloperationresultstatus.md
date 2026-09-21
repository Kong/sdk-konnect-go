# DataPlaneNodeLogLevelOperationResultStatus

Per-node status for one targeted Data Plane node.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DataPlaneNodeLogLevelOperationResultStatusInProgress

// Open enum: custom values can be created with a direct type cast
custom := components.DataPlaneNodeLogLevelOperationResultStatus("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `DataPlaneNodeLogLevelOperationResultStatusInProgress`  | in_progress                                             |
| `DataPlaneNodeLogLevelOperationResultStatusApplied`     | applied                                                 |
| `DataPlaneNodeLogLevelOperationResultStatusReverted`    | reverted                                                |
| `DataPlaneNodeLogLevelOperationResultStatusSuperseded`  | superseded                                              |
| `DataPlaneNodeLogLevelOperationResultStatusFailed`      | failed                                                  |
| `DataPlaneNodeLogLevelOperationResultStatusUnsupported` | unsupported                                             |