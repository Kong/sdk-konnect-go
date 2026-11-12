# AuthType

The auth type value of the cluster associated with the Runtime Group.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AuthTypePinnedClientCerts

// Open enum: custom values can be created with a direct type cast
custom := components.AuthType("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `AuthTypePinnedClientCerts` | pinned_client_certs         |
| `AuthTypePkiClientCerts`    | pki_client_certs            |