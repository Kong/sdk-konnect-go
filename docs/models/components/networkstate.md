# NetworkState

The current state of the network. Possible values:
- `created`: The network was created, but provisioning hasn't started yet.
- `initializing`: The network is being provisioned in the cloud provider.
- `offline`: The network was created but isn't provisioned; no data planes can be deployed.
- `ready`: The network is fully provisioned and available for data plane groups.
- `terminating`: The network is being deleted and is no longer accepting new resources.
- `terminated`: The network has been fully deleted and is no longer available.


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