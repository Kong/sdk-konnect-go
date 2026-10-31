# AuthServerTokenEndpointAuthMethod

Requested authentication method for OAuth 2.0 endpoints.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AuthServerTokenEndpointAuthMethodPrivateKeyJwt

// Open enum: custom values can be created with a direct type cast
custom := components.AuthServerTokenEndpointAuthMethod("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `AuthServerTokenEndpointAuthMethodPrivateKeyJwt`    | private_key_jwt                                     |
| `AuthServerTokenEndpointAuthMethodClientSecretPost` | client_secret_post                                  |
| `AuthServerTokenEndpointAuthMethodNone`             | none                                                |