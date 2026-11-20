# AIGatewayIdentityProviderOpenIDConnectAuthorizationCookieSameSite

Controls whether a cookie is sent with cross-origin requests, providing some protection against cross-site request forgery attacks.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectAuthorizationCookieSameSiteDefault

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectAuthorizationCookieSameSite("custom_value")
```


## Values

| Name                                                                       | Value                                                                      |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectAuthorizationCookieSameSiteDefault` | Default                                                                    |
| `AIGatewayIdentityProviderOpenIDConnectAuthorizationCookieSameSiteLax`     | Lax                                                                        |
| `AIGatewayIdentityProviderOpenIDConnectAuthorizationCookieSameSiteNone`    | None                                                                       |
| `AIGatewayIdentityProviderOpenIDConnectAuthorizationCookieSameSiteStrict`  | Strict                                                                     |