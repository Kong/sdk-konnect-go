# AIGatewayAuthStrategyOpenIDConnectClientAuth

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectClientAuthClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectClientAuth("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectClientAuthClientSecretBasic`       | client_secret_basic                                                   |
| `AIGatewayAuthStrategyOpenIDConnectClientAuthClientSecretJwt`         | client_secret_jwt                                                     |
| `AIGatewayAuthStrategyOpenIDConnectClientAuthClientSecretPost`        | client_secret_post                                                    |
| `AIGatewayAuthStrategyOpenIDConnectClientAuthNone`                    | none                                                                  |
| `AIGatewayAuthStrategyOpenIDConnectClientAuthPrivateKeyJwt`           | private_key_jwt                                                       |
| `AIGatewayAuthStrategyOpenIDConnectClientAuthSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                           |
| `AIGatewayAuthStrategyOpenIDConnectClientAuthTLSClientAuth`           | tls_client_auth                                                       |