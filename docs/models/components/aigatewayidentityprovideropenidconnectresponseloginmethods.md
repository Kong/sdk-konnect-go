# AIGatewayIdentityProviderOpenIDConnectResponseLoginMethods

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectResponseLoginMethodsAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectResponseLoginMethods("custom_value")
```


## Values

| Name                                                                          | Value                                                                         |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectResponseLoginMethodsAuthorizationCode` | authorization_code                                                            |
| `AIGatewayIdentityProviderOpenIDConnectResponseLoginMethodsBearer`            | bearer                                                                        |
| `AIGatewayIdentityProviderOpenIDConnectResponseLoginMethodsClientCredentials` | client_credentials                                                            |
| `AIGatewayIdentityProviderOpenIDConnectResponseLoginMethodsIntrospection`     | introspection                                                                 |
| `AIGatewayIdentityProviderOpenIDConnectResponseLoginMethodsKongOauth2`        | kong_oauth2                                                                   |
| `AIGatewayIdentityProviderOpenIDConnectResponseLoginMethodsPassword`          | password                                                                      |
| `AIGatewayIdentityProviderOpenIDConnectResponseLoginMethodsRefreshToken`      | refresh_token                                                                 |
| `AIGatewayIdentityProviderOpenIDConnectResponseLoginMethodsSession`           | session                                                                       |
| `AIGatewayIdentityProviderOpenIDConnectResponseLoginMethodsUserinfo`          | userinfo                                                                      |