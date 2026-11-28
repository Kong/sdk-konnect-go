# AIGatewayIdentityProviderOpenIDConnectResponseLoginRedirectMode

Where to place `login_tokens` when using `redirect` `login_action`: - `query`: place tokens in query string - `fragment`: place tokens in url fragment (not readable by servers).

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectResponseLoginRedirectModeFragment

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectResponseLoginRedirectMode("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectResponseLoginRedirectModeFragment` | fragment                                                                  |
| `AIGatewayIdentityProviderOpenIDConnectResponseLoginRedirectModeQuery`    | query                                                                     |