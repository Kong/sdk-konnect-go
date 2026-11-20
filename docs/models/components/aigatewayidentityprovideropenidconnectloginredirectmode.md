# AIGatewayIdentityProviderOpenIDConnectLoginRedirectMode

Where to place `login_tokens` when using `redirect` `login_action`: - `query`: place tokens in query string - `fragment`: place tokens in url fragment (not readable by servers).

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectLoginRedirectModeFragment

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectLoginRedirectMode("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectLoginRedirectModeFragment` | fragment                                                          |
| `AIGatewayIdentityProviderOpenIDConnectLoginRedirectModeQuery`    | query                                                             |