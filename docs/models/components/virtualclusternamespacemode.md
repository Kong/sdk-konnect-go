# VirtualClusterNamespaceMode

* hide_prefix - the configured prefix is hidden from clients for topics and IDs when reading.


  Created resources are written with the prefix on the backend cluster.
* enforce_prefix - the configured prefix remains visible to clients.


  Created resources must include the prefix or the request will fail.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.VirtualClusterNamespaceModeHidePrefix

// Open enum: custom values can be created with a direct type cast
custom := components.VirtualClusterNamespaceMode("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `VirtualClusterNamespaceModeHidePrefix`    | hide_prefix                                |
| `VirtualClusterNamespaceModeEnforcePrefix` | enforce_prefix                             |