# AIGatewayIdentityProviderOpenIDConnectResponseProvider

Identity provider used with `grant_type = jwt_bearer`; not allowed with `grant_type = token_exchange`. `standard` (default) adds no provider-specific parameter. `microsoft` adds the required `requested_token_use=on_behalf_of`, which is fixed and cannot be overridden via `post_args_names/values`.

**Requires a minimum runtime version of `2.3`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectResponseProviderMicrosoft

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectResponseProvider("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectResponseProviderMicrosoft` | microsoft                                                         |
| `AIGatewayIdentityProviderOpenIDConnectResponseProviderStandard`  | standard                                                          |