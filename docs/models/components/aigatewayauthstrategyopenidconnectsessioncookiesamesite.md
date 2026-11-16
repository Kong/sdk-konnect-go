# AIGatewayAuthStrategyOpenIDConnectSessionCookieSameSite

Controls whether a cookie is sent with cross-origin requests, providing some protection against cross-site request forgery attacks.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectSessionCookieSameSiteDefault

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectSessionCookieSameSite("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectSessionCookieSameSiteDefault` | Default                                                          |
| `AIGatewayAuthStrategyOpenIDConnectSessionCookieSameSiteLax`     | Lax                                                              |
| `AIGatewayAuthStrategyOpenIDConnectSessionCookieSameSiteNone`    | None                                                             |
| `AIGatewayAuthStrategyOpenIDConnectSessionCookieSameSiteStrict`  | Strict                                                           |