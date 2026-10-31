# ProvisionedBy

Identifies the actor that provisioned this principal. Defaults to `konnect_user` for principals created directly by a user; other values identify a service that provisioned the principal on a user's behalf (for example, `developer_portal` or `trusted_idp`).

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ProvisionedByKonnectUser

// Open enum: custom values can be created with a direct type cast
custom := components.ProvisionedBy("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `ProvisionedByKonnectUser`     | konnect_user                   |
| `ProvisionedByDeveloperPortal` | developer_portal               |
| `ProvisionedByTrustedIdp`      | trusted_idp                    |