# AIGatewayAuthStrategyOpenIDConnectTokenHeadersGrants

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectTokenHeadersGrantsAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectTokenHeadersGrants("custom_value")
```


## Values

| Name                                                                    | Value                                                                   |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectTokenHeadersGrantsAuthorizationCode` | authorization_code                                                      |
| `AIGatewayAuthStrategyOpenIDConnectTokenHeadersGrantsClientCredentials` | client_credentials                                                      |
| `AIGatewayAuthStrategyOpenIDConnectTokenHeadersGrantsPassword`          | password                                                                |
| `AIGatewayAuthStrategyOpenIDConnectTokenHeadersGrantsRefreshToken`      | refresh_token                                                           |