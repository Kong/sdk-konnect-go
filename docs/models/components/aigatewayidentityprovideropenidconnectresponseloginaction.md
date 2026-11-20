# AIGatewayIdentityProviderOpenIDConnectResponseLoginAction

What to do after successful login: - `upstream`: proxy request to upstream service - `response`: terminate request with a response - `redirect`: redirect to a different location.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectResponseLoginActionRedirect

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectResponseLoginAction("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectResponseLoginActionRedirect` | redirect                                                            |
| `AIGatewayIdentityProviderOpenIDConnectResponseLoginActionResponse` | response                                                            |
| `AIGatewayIdentityProviderOpenIDConnectResponseLoginActionUpstream` | upstream                                                            |