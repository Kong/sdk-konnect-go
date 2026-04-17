# GroupStatusState

The state of the control plane group.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.GroupStatusStateOk

// Open enum: custom values can be created with a direct type cast
custom := components.GroupStatusState("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `GroupStatusStateOk`       | OK                         |
| `GroupStatusStateConflict` | CONFLICT                   |
| `GroupStatusStateUnknown`  | UNKNOWN                    |