# TokenEndpointAuthMethod

The token endpoint authentication method: `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.TokenEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.TokenEndpointAuthMethod("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `TokenEndpointAuthMethodClientSecretBasic`       | client_secret_basic                              |
| `TokenEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                |
| `TokenEndpointAuthMethodClientSecretPost`        | client_secret_post                               |
| `TokenEndpointAuthMethodNone`                    | none                                             |
| `TokenEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                  |
| `TokenEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                      |
| `TokenEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                  |