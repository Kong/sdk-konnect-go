# AuthStrategyClientCredentialsCredentialType

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AuthStrategyClientCredentialsCredentialTypeClientCredentials

// Open enum: custom values can be created with a direct type cast
custom := components.AuthStrategyClientCredentialsCredentialType("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `AuthStrategyClientCredentialsCredentialTypeClientCredentials`            | client_credentials                                                        |
| `AuthStrategyClientCredentialsCredentialTypeSelfManagedClientCredentials` | self_managed_client_credentials                                           |