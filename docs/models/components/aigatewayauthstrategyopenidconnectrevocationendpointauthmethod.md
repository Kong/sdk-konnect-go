# AIGatewayAuthStrategyOpenIDConnectRevocationEndpointAuthMethod

The revocation endpoint authentication method: : `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectRevocationEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectRevocationEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                                    | Value                                                                                   |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectRevocationEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                     |
| `AIGatewayAuthStrategyOpenIDConnectRevocationEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                       |
| `AIGatewayAuthStrategyOpenIDConnectRevocationEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                      |
| `AIGatewayAuthStrategyOpenIDConnectRevocationEndpointAuthMethodNone`                    | none                                                                                    |
| `AIGatewayAuthStrategyOpenIDConnectRevocationEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                         |
| `AIGatewayAuthStrategyOpenIDConnectRevocationEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                             |
| `AIGatewayAuthStrategyOpenIDConnectRevocationEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                         |