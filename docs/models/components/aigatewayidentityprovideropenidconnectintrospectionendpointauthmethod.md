# AIGatewayIdentityProviderOpenIDConnectIntrospectionEndpointAuthMethod

The introspection endpoint authentication method: : `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectIntrospectionEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectIntrospectionEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                                           | Value                                                                                          |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectIntrospectionEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                            |
| `AIGatewayIdentityProviderOpenIDConnectIntrospectionEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                              |
| `AIGatewayIdentityProviderOpenIDConnectIntrospectionEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                             |
| `AIGatewayIdentityProviderOpenIDConnectIntrospectionEndpointAuthMethodNone`                    | none                                                                                           |
| `AIGatewayIdentityProviderOpenIDConnectIntrospectionEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                                |
| `AIGatewayIdentityProviderOpenIDConnectIntrospectionEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                                    |
| `AIGatewayIdentityProviderOpenIDConnectIntrospectionEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                                |