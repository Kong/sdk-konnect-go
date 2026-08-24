# AIGatewayAuthStrategyOpenIDConnectAuthMethods

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectAuthMethodsAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectAuthMethods("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectAuthMethodsAuthorizationCode` | authorization_code                                               |
| `AIGatewayAuthStrategyOpenIDConnectAuthMethodsBearer`            | bearer                                                           |
| `AIGatewayAuthStrategyOpenIDConnectAuthMethodsClientCredentials` | client_credentials                                               |
| `AIGatewayAuthStrategyOpenIDConnectAuthMethodsIntrospection`     | introspection                                                    |
| `AIGatewayAuthStrategyOpenIDConnectAuthMethodsKongOauth2`        | kong_oauth2                                                      |
| `AIGatewayAuthStrategyOpenIDConnectAuthMethodsPassword`          | password                                                         |
| `AIGatewayAuthStrategyOpenIDConnectAuthMethodsRefreshToken`      | refresh_token                                                    |
| `AIGatewayAuthStrategyOpenIDConnectAuthMethodsSession`           | session                                                          |
| `AIGatewayAuthStrategyOpenIDConnectAuthMethodsUserinfo`          | userinfo                                                         |