# AIGatewayIdentityProviderOpenIDConnectResponseRevocationEndpointAuthMethod

The revocation endpoint authentication method: : `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectResponseRevocationEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectResponseRevocationEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                                                | Value                                                                                               |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectResponseRevocationEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                                 |
| `AIGatewayIdentityProviderOpenIDConnectResponseRevocationEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                                   |
| `AIGatewayIdentityProviderOpenIDConnectResponseRevocationEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                                  |
| `AIGatewayIdentityProviderOpenIDConnectResponseRevocationEndpointAuthMethodNone`                    | none                                                                                                |
| `AIGatewayIdentityProviderOpenIDConnectResponseRevocationEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                                     |
| `AIGatewayIdentityProviderOpenIDConnectResponseRevocationEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                                         |
| `AIGatewayIdentityProviderOpenIDConnectResponseRevocationEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                                     |