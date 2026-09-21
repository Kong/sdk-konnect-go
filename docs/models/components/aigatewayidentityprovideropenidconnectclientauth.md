# AIGatewayIdentityProviderOpenIDConnectClientAuth

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectClientAuthClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectClientAuth("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectClientAuthClientSecretBasic`       | client_secret_basic                                                       |
| `AIGatewayIdentityProviderOpenIDConnectClientAuthClientSecretPost`        | client_secret_post                                                        |
| `AIGatewayIdentityProviderOpenIDConnectClientAuthClientSecretJwt`         | client_secret_jwt                                                         |
| `AIGatewayIdentityProviderOpenIDConnectClientAuthPrivateKeyJwt`           | private_key_jwt                                                           |
| `AIGatewayIdentityProviderOpenIDConnectClientAuthTLSClientAuth`           | tls_client_auth                                                           |
| `AIGatewayIdentityProviderOpenIDConnectClientAuthSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                               |
| `AIGatewayIdentityProviderOpenIDConnectClientAuthNone`                    | none                                                                      |