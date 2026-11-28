# AIGatewayIdentityProviderOpenIDConnectResponseResponseMode

Response mode passed to the authorization endpoint: - `query`: for parameters in query string - `form_post`: for parameters in request body - `fragment`: for parameters in uri fragment (rarely useful as the plugin itself cannot read it) - `query.jwt`, `form_post.jwt`, `fragment.jwt`: similar to `query`, `form_post` and `fragment` but the parameters are encoded in a JWT - `jwt`: shortcut that indicates the default encoding for the requested response type.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectResponseResponseModeFormPost

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectResponseResponseMode("custom_value")
```


## Values

| Name                                                                    | Value                                                                   |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectResponseResponseModeFormPost`    | form_post                                                               |
| `AIGatewayIdentityProviderOpenIDConnectResponseResponseModeFormPostJwt` | form_post.jwt                                                           |
| `AIGatewayIdentityProviderOpenIDConnectResponseResponseModeFragment`    | fragment                                                                |
| `AIGatewayIdentityProviderOpenIDConnectResponseResponseModeFragmentJwt` | fragment.jwt                                                            |
| `AIGatewayIdentityProviderOpenIDConnectResponseResponseModeJwt`         | jwt                                                                     |
| `AIGatewayIdentityProviderOpenIDConnectResponseResponseModeQuery`       | query                                                                   |
| `AIGatewayIdentityProviderOpenIDConnectResponseResponseModeQueryJwt`    | query.jwt                                                               |