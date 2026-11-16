# AIGatewayAuthStrategyOpenIDConnectPushedAuthorizationRequestEndpointAuthMethod

The pushed authorization request endpoint authentication method: `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectPushedAuthorizationRequestEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                                                    | Value                                                                                                   |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                                     |
| `AIGatewayAuthStrategyOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                                       |
| `AIGatewayAuthStrategyOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                                      |
| `AIGatewayAuthStrategyOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodNone`                    | none                                                                                                    |
| `AIGatewayAuthStrategyOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                                         |
| `AIGatewayAuthStrategyOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                                             |
| `AIGatewayAuthStrategyOpenIDConnectPushedAuthorizationRequestEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                                         |