# UpdateControlPlaneRequestAuthType

The auth type value of the cluster associated with the Runtime Group.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.UpdateControlPlaneRequestAuthTypePinnedClientCerts

// Open enum: custom values can be created with a direct type cast
custom := components.UpdateControlPlaneRequestAuthType("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `UpdateControlPlaneRequestAuthTypePinnedClientCerts` | pinned_client_certs                                  |
| `UpdateControlPlaneRequestAuthTypePkiClientCerts`    | pki_client_certs                                     |