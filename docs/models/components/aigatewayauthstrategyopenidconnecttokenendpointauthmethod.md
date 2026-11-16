# AIGatewayAuthStrategyOpenIDConnectTokenEndpointAuthMethod

The token endpoint authentication method: `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectTokenEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectTokenEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                               | Value                                                                              |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectTokenEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                |
| `AIGatewayAuthStrategyOpenIDConnectTokenEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                  |
| `AIGatewayAuthStrategyOpenIDConnectTokenEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                 |
| `AIGatewayAuthStrategyOpenIDConnectTokenEndpointAuthMethodNone`                    | none                                                                               |
| `AIGatewayAuthStrategyOpenIDConnectTokenEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                    |
| `AIGatewayAuthStrategyOpenIDConnectTokenEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                        |
| `AIGatewayAuthStrategyOpenIDConnectTokenEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                    |