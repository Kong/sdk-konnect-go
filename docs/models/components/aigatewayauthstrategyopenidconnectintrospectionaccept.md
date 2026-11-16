# AIGatewayAuthStrategyOpenIDConnectIntrospectionAccept

The value of `Accept` header for introspection requests: - `application/json`: introspection response as JSON - `application/token-introspection+jwt`: introspection response as JWT (from the current IETF draft document) - `application/jwt`: introspection response as JWT (from the obsolete IETF draft document).

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectIntrospectionAcceptApplicationJSON

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectIntrospectionAccept("custom_value")
```


## Values

| Name                                                                                        | Value                                                                                       |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectIntrospectionAcceptApplicationJSON`                      | application/json                                                                            |
| `AIGatewayAuthStrategyOpenIDConnectIntrospectionAcceptApplicationJwt`                       | application/jwt                                                                             |
| `AIGatewayAuthStrategyOpenIDConnectIntrospectionAcceptApplicationTokenIntrospectionPlusJwt` | application/token-introspection+jwt                                                         |