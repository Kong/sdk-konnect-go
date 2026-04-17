# AppAuthStrategyOpenIDConnectResponseProviderType

The type of DCR provider.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AppAuthStrategyOpenIDConnectResponseProviderTypeAuth0

// Open enum: custom values can be created with a direct type cast
custom := components.AppAuthStrategyOpenIDConnectResponseProviderType("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `AppAuthStrategyOpenIDConnectResponseProviderTypeAuth0`   | auth0                                                     |
| `AppAuthStrategyOpenIDConnectResponseProviderTypeAzureAd` | azureAd                                                   |
| `AppAuthStrategyOpenIDConnectResponseProviderTypeCurity`  | curity                                                    |
| `AppAuthStrategyOpenIDConnectResponseProviderTypeOkta`    | okta                                                      |
| `AppAuthStrategyOpenIDConnectResponseProviderTypeHTTP`    | http                                                      |