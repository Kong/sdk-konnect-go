# AIGatewayIdentityProviderOpenIDConnectPushedAuthorizationRequestEndpointAuthMethod

The pushed authorization request endpoint authentication method: `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectPushedAuthorizationRequestEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                                                        | Value                                                                                                       |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                                         |
| `AIGatewayIdentityProviderOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                                           |
| `AIGatewayIdentityProviderOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                                          |
| `AIGatewayIdentityProviderOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodNone`                    | none                                                                                                        |
| `AIGatewayIdentityProviderOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                                             |
| `AIGatewayIdentityProviderOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                                                 |
| `AIGatewayIdentityProviderOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                                             |