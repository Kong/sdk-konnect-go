# RevocationEndpointAuthMethod

The revocation endpoint authentication method: : `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.RevocationEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.RevocationEndpointAuthMethod("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `RevocationEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                   |
| `RevocationEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                     |
| `RevocationEndpointAuthMethodClientSecretPost`        | client_secret_post                                    |
| `RevocationEndpointAuthMethodNone`                    | none                                                  |
| `RevocationEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                       |
| `RevocationEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                           |
| `RevocationEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                       |