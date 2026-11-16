# AIGatewayAuthStrategyOpenIDConnectResponseTokenEndpointAuthMethod

The token endpoint authentication method: `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectResponseTokenEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectResponseTokenEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                                       | Value                                                                                      |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `AIGatewayAuthStrategyOpenIDConnectResponseTokenEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                        |
| `AIGatewayAuthStrategyOpenIDConnectResponseTokenEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                          |
| `AIGatewayAuthStrategyOpenIDConnectResponseTokenEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                         |
| `AIGatewayAuthStrategyOpenIDConnectResponseTokenEndpointAuthMethodNone`                    | none                                                                                       |
| `AIGatewayAuthStrategyOpenIDConnectResponseTokenEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                            |
| `AIGatewayAuthStrategyOpenIDConnectResponseTokenEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                                |
| `AIGatewayAuthStrategyOpenIDConnectResponseTokenEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                            |