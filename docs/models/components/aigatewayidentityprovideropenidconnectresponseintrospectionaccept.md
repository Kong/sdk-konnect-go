# AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionAccept

The value of `Accept` header for introspection requests: - `application/json`: introspection response as JSON - `application/token-introspection+jwt`: introspection response as JWT (from the current IETF draft document) - `application/jwt`: introspection response as JWT (from the obsolete IETF draft document).

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionAcceptApplicationJSON

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionAccept("custom_value")
```


## Values

| Name                                                                                                    | Value                                                                                                   |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionAcceptApplicationJSON`                      | application/json                                                                                        |
| `AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionAcceptApplicationJwt`                       | application/jwt                                                                                         |
| `AIGatewayIdentityProviderOpenIDConnectResponseIntrospectionAcceptApplicationTokenIntrospectionPlusJwt` | application/token-introspection+jwt                                                                     |