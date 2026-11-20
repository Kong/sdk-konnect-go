# AIGatewayIdentityProviderOpenIDConnectTokenEndpointAuthMethod

The token endpoint authentication method: `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectTokenEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectTokenEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                                   | Value                                                                                  |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectTokenEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                    |
| `AIGatewayIdentityProviderOpenIDConnectTokenEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                      |
| `AIGatewayIdentityProviderOpenIDConnectTokenEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                     |
| `AIGatewayIdentityProviderOpenIDConnectTokenEndpointAuthMethodNone`                    | none                                                                                   |
| `AIGatewayIdentityProviderOpenIDConnectTokenEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                        |
| `AIGatewayIdentityProviderOpenIDConnectTokenEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                            |
| `AIGatewayIdentityProviderOpenIDConnectTokenEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                        |