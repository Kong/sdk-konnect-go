# AIGatewayIdentityProviderOpenIDConnectResponseAuthMethods

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectResponseAuthMethodsAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectResponseAuthMethods("custom_value")
```


## Values

| Name                                                                         | Value                                                                        |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectResponseAuthMethodsAuthorizationCode` | authorization_code                                                           |
| `AIGatewayIdentityProviderOpenIDConnectResponseAuthMethodsBearer`            | bearer                                                                       |
| `AIGatewayIdentityProviderOpenIDConnectResponseAuthMethodsClientCredentials` | client_credentials                                                           |
| `AIGatewayIdentityProviderOpenIDConnectResponseAuthMethodsIntrospection`     | introspection                                                                |
| `AIGatewayIdentityProviderOpenIDConnectResponseAuthMethodsKongOauth2`        | kong_oauth2                                                                  |
| `AIGatewayIdentityProviderOpenIDConnectResponseAuthMethodsPassword`          | password                                                                     |
| `AIGatewayIdentityProviderOpenIDConnectResponseAuthMethodsRefreshToken`      | refresh_token                                                                |
| `AIGatewayIdentityProviderOpenIDConnectResponseAuthMethodsSession`           | session                                                                      |
| `AIGatewayIdentityProviderOpenIDConnectResponseAuthMethodsUserinfo`          | userinfo                                                                     |