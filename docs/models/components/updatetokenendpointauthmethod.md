# UpdateTokenEndpointAuthMethod

Requested authentication method for OAuth 2.0 endpoints.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.UpdateTokenEndpointAuthMethodPrivateKeyJwt

// Open enum: custom values can be created with a direct type cast
custom := components.UpdateTokenEndpointAuthMethod("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `UpdateTokenEndpointAuthMethodPrivateKeyJwt`    | private_key_jwt                                 |
| `UpdateTokenEndpointAuthMethodClientSecretPost` | client_secret_post                              |
| `UpdateTokenEndpointAuthMethodNone`             | none                                            |