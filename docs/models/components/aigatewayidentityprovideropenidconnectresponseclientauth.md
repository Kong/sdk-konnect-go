# AIGatewayIdentityProviderOpenIDConnectResponseClientAuth

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectResponseClientAuthClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectResponseClientAuth("custom_value")
```


## Values

| Name                                                                              | Value                                                                             |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectResponseClientAuthClientSecretBasic`       | client_secret_basic                                                               |
| `AIGatewayIdentityProviderOpenIDConnectResponseClientAuthClientSecretPost`        | client_secret_post                                                                |
| `AIGatewayIdentityProviderOpenIDConnectResponseClientAuthClientSecretJwt`         | client_secret_jwt                                                                 |
| `AIGatewayIdentityProviderOpenIDConnectResponseClientAuthPrivateKeyJwt`           | private_key_jwt                                                                   |
| `AIGatewayIdentityProviderOpenIDConnectResponseClientAuthTLSClientAuth`           | tls_client_auth                                                                   |
| `AIGatewayIdentityProviderOpenIDConnectResponseClientAuthSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                       |
| `AIGatewayIdentityProviderOpenIDConnectResponseClientAuthNone`                    | none                                                                              |