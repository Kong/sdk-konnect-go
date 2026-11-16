# AIGatewayAuthStrategyOpenIDConnectLoginAction

What to do after successful login: - `upstream`: proxy request to upstream service - `response`: terminate request with a response - `redirect`: redirect to a different location.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectLoginActionRedirect

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectLoginAction("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectLoginActionRedirect` | redirect                                                |
| `AIGatewayAuthStrategyOpenIDConnectLoginActionResponse` | response                                                |
| `AIGatewayAuthStrategyOpenIDConnectLoginActionUpstream` | upstream                                                |