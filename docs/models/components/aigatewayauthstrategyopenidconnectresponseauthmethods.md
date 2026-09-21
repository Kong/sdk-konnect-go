# AIGatewayAuthStrategyOpenIDConnectResponseAuthMethods

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectResponseAuthMethodsAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectResponseAuthMethods("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `AIGatewayAuthStrategyOpenIDConnectResponseAuthMethodsAuthorizationCode` | authorization_code                                                       |
| `AIGatewayAuthStrategyOpenIDConnectResponseAuthMethodsBearer`            | bearer                                                                   |
| `AIGatewayAuthStrategyOpenIDConnectResponseAuthMethodsClientCredentials` | client_credentials                                                       |
| `AIGatewayAuthStrategyOpenIDConnectResponseAuthMethodsIntrospection`     | introspection                                                            |
| `AIGatewayAuthStrategyOpenIDConnectResponseAuthMethodsKongOauth2`        | kong_oauth2                                                              |
| `AIGatewayAuthStrategyOpenIDConnectResponseAuthMethodsPassword`          | password                                                                 |
| `AIGatewayAuthStrategyOpenIDConnectResponseAuthMethodsRefreshToken`      | refresh_token                                                            |
| `AIGatewayAuthStrategyOpenIDConnectResponseAuthMethodsSession`           | session                                                                  |
| `AIGatewayAuthStrategyOpenIDConnectResponseAuthMethodsUserinfo`          | userinfo                                                                 |