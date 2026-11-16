# AIGatewayAuthStrategyOpenIDConnectDisableSession

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectDisableSessionAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectDisableSession("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectDisableSessionAuthorizationCode` | authorization_code                                                  |
| `AIGatewayAuthStrategyOpenIDConnectDisableSessionBearer`            | bearer                                                              |
| `AIGatewayAuthStrategyOpenIDConnectDisableSessionClientCredentials` | client_credentials                                                  |
| `AIGatewayAuthStrategyOpenIDConnectDisableSessionIntrospection`     | introspection                                                       |
| `AIGatewayAuthStrategyOpenIDConnectDisableSessionKongOauth2`        | kong_oauth2                                                         |
| `AIGatewayAuthStrategyOpenIDConnectDisableSessionPassword`          | password                                                            |
| `AIGatewayAuthStrategyOpenIDConnectDisableSessionRefreshToken`      | refresh_token                                                       |
| `AIGatewayAuthStrategyOpenIDConnectDisableSessionSession`           | session                                                             |
| `AIGatewayAuthStrategyOpenIDConnectDisableSessionUserinfo`          | userinfo                                                            |