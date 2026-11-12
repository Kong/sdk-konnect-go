# ControlPlaneAuthType

The auth type value of the cluster associated with the Runtime Group.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ControlPlaneAuthTypePinnedClientCerts

// Open enum: custom values can be created with a direct type cast
custom := components.ControlPlaneAuthType("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `ControlPlaneAuthTypePinnedClientCerts` | pinned_client_certs                     |
| `ControlPlaneAuthTypePkiClientCerts`    | pki_client_certs                        |