# TokenEndpointAuthMethod

Requested authentication method for OAuth 2.0 endpoints.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.TokenEndpointAuthMethodPrivateKeyJwt

// Open enum: custom values can be created with a direct type cast
custom := components.TokenEndpointAuthMethod("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `TokenEndpointAuthMethodPrivateKeyJwt`    | private_key_jwt                           |
| `TokenEndpointAuthMethodClientSecretPost` | client_secret_post                        |
| `TokenEndpointAuthMethodNone`             | none                                      |