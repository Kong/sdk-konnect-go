# AIGatewayAuthStrategyOpenIDConnectIgnoreSignature

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectIgnoreSignatureAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectIgnoreSignature("custom_value")
```


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectIgnoreSignatureAuthorizationCode` | authorization_code                                                   |
| `AIGatewayAuthStrategyOpenIDConnectIgnoreSignatureClientCredentials` | client_credentials                                                   |
| `AIGatewayAuthStrategyOpenIDConnectIgnoreSignatureIntrospection`     | introspection                                                        |
| `AIGatewayAuthStrategyOpenIDConnectIgnoreSignaturePassword`          | password                                                             |
| `AIGatewayAuthStrategyOpenIDConnectIgnoreSignatureRefreshToken`      | refresh_token                                                        |
| `AIGatewayAuthStrategyOpenIDConnectIgnoreSignatureSession`           | session                                                              |
| `AIGatewayAuthStrategyOpenIDConnectIgnoreSignatureUserinfo`          | userinfo                                                             |