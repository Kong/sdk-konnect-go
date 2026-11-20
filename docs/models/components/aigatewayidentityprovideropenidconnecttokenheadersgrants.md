# AIGatewayIdentityProviderOpenIDConnectTokenHeadersGrants

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayIdentityProviderOpenIDConnectTokenHeadersGrantsAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayIdentityProviderOpenIDConnectTokenHeadersGrants("custom_value")
```


## Values

| Name                                                                        | Value                                                                       |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `AIGatewayIdentityProviderOpenIDConnectTokenHeadersGrantsAuthorizationCode` | authorization_code                                                          |
| `AIGatewayIdentityProviderOpenIDConnectTokenHeadersGrantsClientCredentials` | client_credentials                                                          |
| `AIGatewayIdentityProviderOpenIDConnectTokenHeadersGrantsPassword`          | password                                                                    |
| `AIGatewayIdentityProviderOpenIDConnectTokenHeadersGrantsRefreshToken`      | refresh_token                                                               |