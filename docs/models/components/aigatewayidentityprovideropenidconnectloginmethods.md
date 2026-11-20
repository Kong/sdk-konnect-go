# AIGatewayIdentityProviderOpenIDConnectLoginMethods

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectLoginMethodsAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectLoginMethods("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectLoginMethodsAuthorizationCode` | authorization_code                                                    |
| `AIGatewayIdentityProviderOpenIDConnectLoginMethodsBearer`            | bearer                                                                |
| `AIGatewayIdentityProviderOpenIDConnectLoginMethodsClientCredentials` | client_credentials                                                    |
| `AIGatewayIdentityProviderOpenIDConnectLoginMethodsIntrospection`     | introspection                                                         |
| `AIGatewayIdentityProviderOpenIDConnectLoginMethodsKongOauth2`        | kong_oauth2                                                           |
| `AIGatewayIdentityProviderOpenIDConnectLoginMethodsPassword`          | password                                                              |
| `AIGatewayIdentityProviderOpenIDConnectLoginMethodsRefreshToken`      | refresh_token                                                         |
| `AIGatewayIdentityProviderOpenIDConnectLoginMethodsSession`           | session                                                               |
| `AIGatewayIdentityProviderOpenIDConnectLoginMethodsUserinfo`          | userinfo                                                              |