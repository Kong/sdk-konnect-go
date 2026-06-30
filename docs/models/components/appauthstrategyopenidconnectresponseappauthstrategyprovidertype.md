# AppAuthStrategyOpenIDConnectResponseAppAuthStrategyProviderType

The type of DCR provider.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyProviderTypeAuth0

// Open enum: custom values can be created with a direct type cast
custom := components.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyProviderType("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `AppAuthStrategyOpenIDConnectResponseAppAuthStrategyProviderTypeAuth0`   | auth0                                                                    |
| `AppAuthStrategyOpenIDConnectResponseAppAuthStrategyProviderTypeAzureAd` | azureAd                                                                  |
| `AppAuthStrategyOpenIDConnectResponseAppAuthStrategyProviderTypeCurity`  | curity                                                                   |
| `AppAuthStrategyOpenIDConnectResponseAppAuthStrategyProviderTypeOkta`    | okta                                                                     |
| `AppAuthStrategyOpenIDConnectResponseAppAuthStrategyProviderTypeHTTP`    | http                                                                     |