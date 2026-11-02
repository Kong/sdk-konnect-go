# AIGWOpenIDConnectGeneratedConfigGrantType

The token exchange grant. `token_exchange` (default) uses the OAuth 2.0 Token Exchange grant (RFC 8693). `jwt_bearer` uses the JWT Bearer authorization-grant flow (RFC 7523); see `token_exchange.provider` for provider-specific defaults in this mode.

**Requires a minimum runtime version of `2.3`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGWOpenIDConnectGeneratedConfigGrantTypeJwtBearer

// Open enum: custom values can be created with a direct type cast
custom := components.AIGWOpenIDConnectGeneratedConfigGrantType("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `AIGWOpenIDConnectGeneratedConfigGrantTypeJwtBearer`     | jwt_bearer                                               |
| `AIGWOpenIDConnectGeneratedConfigGrantTypeTokenExchange` | token_exchange                                           |