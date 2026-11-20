# AIGatewayIdentityProviderOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethod

The pushed authorization request endpoint authentication method: `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                                                                | Value                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                                                 |
| `AIGatewayIdentityProviderOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                                                   |
| `AIGatewayIdentityProviderOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                                                  |
| `AIGatewayIdentityProviderOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodNone`                    | none                                                                                                                |
| `AIGatewayIdentityProviderOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                                                     |
| `AIGatewayIdentityProviderOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                                                         |
| `AIGatewayIdentityProviderOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                                                     |