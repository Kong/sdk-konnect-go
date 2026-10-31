# AddOnState

The current state of the add-on. Possible values:
- `initializing` - The add-on is in the process of being initialized/updated.
- `ready` - The add-on is fully operational.
- `terminating` - The add-on is in the process of being deleted.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AddOnStateInitializing

// Open enum: custom values can be created with a direct type cast
custom := components.AddOnState("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `AddOnStateInitializing` | initializing             |
| `AddOnStateReady`        | ready                    |
| `AddOnStateTerminating`  | terminating              |