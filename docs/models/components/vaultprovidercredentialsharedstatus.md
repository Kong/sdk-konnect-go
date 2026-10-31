# VaultProviderCredentialSharedStatus

"pending" until admin enrollment completes; "active" once the secret exists.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.VaultProviderCredentialSharedStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.VaultProviderCredentialSharedStatus("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `VaultProviderCredentialSharedStatusPending` | pending                                      |
| `VaultProviderCredentialSharedStatusActive`  | active                                       |