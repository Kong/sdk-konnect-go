# AIGatewayIdentityProviderOpenIDConnectDisableSession

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectDisableSessionAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectDisableSession("custom_value")
```


## Values

| Name                                                                    | Value                                                                   |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectDisableSessionAuthorizationCode` | authorization_code                                                      |
| `AIGatewayIdentityProviderOpenIDConnectDisableSessionBearer`            | bearer                                                                  |
| `AIGatewayIdentityProviderOpenIDConnectDisableSessionClientCredentials` | client_credentials                                                      |
| `AIGatewayIdentityProviderOpenIDConnectDisableSessionIntrospection`     | introspection                                                           |
| `AIGatewayIdentityProviderOpenIDConnectDisableSessionKongOauth2`        | kong_oauth2                                                             |
| `AIGatewayIdentityProviderOpenIDConnectDisableSessionPassword`          | password                                                                |
| `AIGatewayIdentityProviderOpenIDConnectDisableSessionRefreshToken`      | refresh_token                                                           |
| `AIGatewayIdentityProviderOpenIDConnectDisableSessionSession`           | session                                                                 |
| `AIGatewayIdentityProviderOpenIDConnectDisableSessionUserinfo`          | userinfo                                                                |