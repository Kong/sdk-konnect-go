# AIGatewayAuthStrategyOpenIDConnectResponseClientAuth

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectResponseClientAuthClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectResponseClientAuth("custom_value")
```


## Values

| Name                                                                          | Value                                                                         |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectResponseClientAuthClientSecretBasic`       | client_secret_basic                                                           |
| `AIGatewayAuthStrategyOpenIDConnectResponseClientAuthClientSecretPost`        | client_secret_post                                                            |
| `AIGatewayAuthStrategyOpenIDConnectResponseClientAuthClientSecretJwt`         | client_secret_jwt                                                             |
| `AIGatewayAuthStrategyOpenIDConnectResponseClientAuthPrivateKeyJwt`           | private_key_jwt                                                               |
| `AIGatewayAuthStrategyOpenIDConnectResponseClientAuthTLSClientAuth`           | tls_client_auth                                                               |
| `AIGatewayAuthStrategyOpenIDConnectResponseClientAuthSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                   |
| `AIGatewayAuthStrategyOpenIDConnectResponseClientAuthNone`                    | none                                                                          |