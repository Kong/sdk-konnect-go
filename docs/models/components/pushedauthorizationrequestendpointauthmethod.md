# PushedAuthorizationRequestEndpointAuthMethod

The pushed authorization request endpoint authentication method: `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PushedAuthorizationRequestEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.PushedAuthorizationRequestEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `PushedAuthorizationRequestEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                   |
| `PushedAuthorizationRequestEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                     |
| `PushedAuthorizationRequestEndpointAuthMethodClientSecretPost`        | client_secret_post                                                    |
| `PushedAuthorizationRequestEndpointAuthMethodNone`                    | none                                                                  |
| `PushedAuthorizationRequestEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                       |
| `PushedAuthorizationRequestEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                           |
| `PushedAuthorizationRequestEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                       |