# IntrospectionEndpointAuthMethod

The introspection endpoint authentication method: : `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.IntrospectionEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.IntrospectionEndpointAuthMethod("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `IntrospectionEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                      |
| `IntrospectionEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                        |
| `IntrospectionEndpointAuthMethodClientSecretPost`        | client_secret_post                                       |
| `IntrospectionEndpointAuthMethodNone`                    | none                                                     |
| `IntrospectionEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                          |
| `IntrospectionEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                              |
| `IntrospectionEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                          |