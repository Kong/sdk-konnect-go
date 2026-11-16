# AIGatewayAuthStrategyOpenIDConnectIntrospectionEndpointAuthMethod

The introspection endpoint authentication method: : `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectIntrospectionEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectIntrospectionEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                                       | Value                                                                                      |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `AIGatewayAuthStrategyOpenIDConnectIntrospectionEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                        |
| `AIGatewayAuthStrategyOpenIDConnectIntrospectionEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                          |
| `AIGatewayAuthStrategyOpenIDConnectIntrospectionEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                         |
| `AIGatewayAuthStrategyOpenIDConnectIntrospectionEndpointAuthMethodNone`                    | none                                                                                       |
| `AIGatewayAuthStrategyOpenIDConnectIntrospectionEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                            |
| `AIGatewayAuthStrategyOpenIDConnectIntrospectionEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                                |
| `AIGatewayAuthStrategyOpenIDConnectIntrospectionEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                            |