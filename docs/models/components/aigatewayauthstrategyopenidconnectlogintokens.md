# AIGatewayAuthStrategyOpenIDConnectLoginTokens

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectLoginTokensAccessToken

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectLoginTokens("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `AIGatewayAuthStrategyOpenIDConnectLoginTokensAccessToken`   | access_token                                                 |
| `AIGatewayAuthStrategyOpenIDConnectLoginTokensIDToken`       | id_token                                                     |
| `AIGatewayAuthStrategyOpenIDConnectLoginTokensIntrospection` | introspection                                                |
| `AIGatewayAuthStrategyOpenIDConnectLoginTokensRefreshToken`  | refresh_token                                                |
| `AIGatewayAuthStrategyOpenIDConnectLoginTokensTokens`        | tokens                                                       |