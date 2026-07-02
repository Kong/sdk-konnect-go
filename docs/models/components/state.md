# State

The state of the control plane group.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.StateOk

// Open enum: custom values can be created with a direct type cast
custom := components.State("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `StateOk`       | OK              |
| `StateConflict` | CONFLICT        |
| `StateUnknown`  | UNKNOWN         |