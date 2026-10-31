# State

Config sync state.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.StateStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := components.State("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `StateStateUnspecified` | STATE_UNSPECIFIED       |
| `StateStateInSync`      | STATE_IN_SYNC           |
| `StateStatePending`     | STATE_PENDING           |
| `StateStateResiliency`  | STATE_RESILIENCY        |