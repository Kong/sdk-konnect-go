# NetworkState

State of the network.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.NetworkStateCreated

// Open enum: custom values can be created with a direct type cast
custom := components.NetworkState("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `NetworkStateCreated`      | created                    |
| `NetworkStateInitializing` | initializing               |
| `NetworkStateOffline`      | offline                    |
| `NetworkStateReady`        | ready                      |
| `NetworkStateTerminating`  | terminating                |
| `NetworkStateTerminated`   | terminated                 |