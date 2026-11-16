# AIGatewayAuthStrategyOpenIDConnectLoginMethods

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectLoginMethodsAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectLoginMethods("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectLoginMethodsAuthorizationCode` | authorization_code                                                |
| `AIGatewayAuthStrategyOpenIDConnectLoginMethodsBearer`            | bearer                                                            |
| `AIGatewayAuthStrategyOpenIDConnectLoginMethodsClientCredentials` | client_credentials                                                |
| `AIGatewayAuthStrategyOpenIDConnectLoginMethodsIntrospection`     | introspection                                                     |
| `AIGatewayAuthStrategyOpenIDConnectLoginMethodsKongOauth2`        | kong_oauth2                                                       |
| `AIGatewayAuthStrategyOpenIDConnectLoginMethodsPassword`          | password                                                          |
| `AIGatewayAuthStrategyOpenIDConnectLoginMethodsRefreshToken`      | refresh_token                                                     |
| `AIGatewayAuthStrategyOpenIDConnectLoginMethodsSession`           | session                                                           |
| `AIGatewayAuthStrategyOpenIDConnectLoginMethodsUserinfo`          | userinfo                                                          |