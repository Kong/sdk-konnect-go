# NetworkCreateState

Initial state for creating a network.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.NetworkCreateStateInitializing

// Open enum: custom values can be created with a direct type cast
custom := components.NetworkCreateState("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `NetworkCreateStateInitializing` | initializing                     |
| `NetworkCreateStateOffline`      | offline                          |