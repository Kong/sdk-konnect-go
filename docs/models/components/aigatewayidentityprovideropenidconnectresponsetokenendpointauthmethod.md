# AIGatewayIdentityProviderOpenIDConnectResponseTokenEndpointAuthMethod

The token endpoint authentication method: `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectResponseTokenEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectResponseTokenEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                                           | Value                                                                                          |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectResponseTokenEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                            |
| `AIGatewayIdentityProviderOpenIDConnectResponseTokenEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                              |
| `AIGatewayIdentityProviderOpenIDConnectResponseTokenEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                             |
| `AIGatewayIdentityProviderOpenIDConnectResponseTokenEndpointAuthMethodNone`                    | none                                                                                           |
| `AIGatewayIdentityProviderOpenIDConnectResponseTokenEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                                |
| `AIGatewayIdentityProviderOpenIDConnectResponseTokenEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                                    |
| `AIGatewayIdentityProviderOpenIDConnectResponseTokenEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                                |