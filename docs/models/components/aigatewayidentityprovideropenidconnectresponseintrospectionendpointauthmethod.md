# AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionEndpointAuthMethod

The introspection endpoint authentication method: : `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                                                   | Value                                                                                                  |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                                    |
| `AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                                      |
| `AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                                     |
| `AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionEndpointAuthMethodNone`                    | none                                                                                                   |
| `AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                                        |
| `AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                                            |
| `AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                                        |