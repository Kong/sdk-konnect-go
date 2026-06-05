# CredentialType

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CredentialTypeClientCredentials

// Open enum: custom values can be created with a direct type cast
custom := components.CredentialType("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `CredentialTypeClientCredentials`            | client_credentials                           |
| `CredentialTypeSelfManagedClientCredentials` | self_managed_client_credentials              |