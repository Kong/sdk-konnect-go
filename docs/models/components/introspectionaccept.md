# IntrospectionAccept

The value of `Accept` header for introspection requests: - `application/json`: introspection response as JSON - `application/token-introspection+jwt`: introspection response as JWT (from the current IETF draft document) - `application/jwt`: introspection response as JWT (from the obsolete IETF draft document).

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.IntrospectionAcceptApplicationJSON

// Open enum: custom values can be created with a direct type cast
custom := components.IntrospectionAccept("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `IntrospectionAcceptApplicationJSON`                      | application/json                                          |
| `IntrospectionAcceptApplicationJwt`                       | application/jwt                                           |
| `IntrospectionAcceptApplicationTokenIntrospectionPlusJwt` | application/token-introspection+jwt                       |