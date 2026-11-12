# AIGatewayIdentityProviderOpenIDConnectAuthMethods

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectAuthMethodsAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectAuthMethods("custom_value")
```


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectAuthMethodsAuthorizationCode` | authorization_code                                                   |
| `AIGatewayIdentityProviderOpenIDConnectAuthMethodsBearer`            | bearer                                                               |
| `AIGatewayIdentityProviderOpenIDConnectAuthMethodsClientCredentials` | client_credentials                                                   |
| `AIGatewayIdentityProviderOpenIDConnectAuthMethodsIntrospection`     | introspection                                                        |
| `AIGatewayIdentityProviderOpenIDConnectAuthMethodsKongOauth2`        | kong_oauth2                                                          |
| `AIGatewayIdentityProviderOpenIDConnectAuthMethodsPassword`          | password                                                             |
| `AIGatewayIdentityProviderOpenIDConnectAuthMethodsRefreshToken`      | refresh_token                                                        |
| `AIGatewayIdentityProviderOpenIDConnectAuthMethodsSession`           | session                                                              |
| `AIGatewayIdentityProviderOpenIDConnectAuthMethodsUserinfo`          | userinfo                                                             |