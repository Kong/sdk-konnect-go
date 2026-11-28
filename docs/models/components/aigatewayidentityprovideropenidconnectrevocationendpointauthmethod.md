# AIGatewayIdentityProviderOpenIDConnectRevocationEndpointAuthMethod

The revocation endpoint authentication method: : `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectRevocationEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectRevocationEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                                        | Value                                                                                       |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectRevocationEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                         |
| `AIGatewayIdentityProviderOpenIDConnectRevocationEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                           |
| `AIGatewayIdentityProviderOpenIDConnectRevocationEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                          |
| `AIGatewayIdentityProviderOpenIDConnectRevocationEndpointAuthMethodNone`                    | none                                                                                        |
| `AIGatewayIdentityProviderOpenIDConnectRevocationEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                             |
| `AIGatewayIdentityProviderOpenIDConnectRevocationEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                                 |
| `AIGatewayIdentityProviderOpenIDConnectRevocationEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                             |