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
| `AIGatewayAuthStrategyOpenIDConnectClientAuthClientSecretPost`        | client_secret_post                                                    |
| `AIGatewayAuthStrategyOpenIDConnectClientAuthClientSecretJwt`         | client_secret_jwt                                                     |
| `AIGatewayAuthStrategyOpenIDConnectClientAuthPrivateKeyJwt`           | private_key_jwt                                                       |
| `AIGatewayAuthStrategyOpenIDConnectClientAuthTLSClientAuth`           | tls_client_auth                                                       |
| `AIGatewayAuthStrategyOpenIDConnectClientAuthSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                           |
| `AIGatewayAuthStrategyOpenIDConnectClientAuthNone`                    | none                                                                  |