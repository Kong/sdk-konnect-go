# AppAuthStrategyKeyAuthResponseCreateAppAuthStrategyResponseProviderType

The type of DCR provider.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AppAuthStrategyKeyAuthResponseCreateAppAuthStrategyResponseProviderTypeAuth0

// Open enum: custom values can be created with a direct type cast
custom := components.AppAuthStrategyKeyAuthResponseCreateAppAuthStrategyResponseProviderType("custom_value")
```


## Values

| Name                                                                             | Value                                                                            |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `AppAuthStrategyKeyAuthResponseCreateAppAuthStrategyResponseProviderTypeAuth0`   | auth0                                                                            |
| `AppAuthStrategyKeyAuthResponseCreateAppAuthStrategyResponseProviderTypeAzureAd` | azureAd                                                                          |
| `AppAuthStrategyKeyAuthResponseCreateAppAuthStrategyResponseProviderTypeCurity`  | curity                                                                           |
| `AppAuthStrategyKeyAuthResponseCreateAppAuthStrategyResponseProviderTypeOkta`    | okta                                                                             |
| `AppAuthStrategyKeyAuthResponseCreateAppAuthStrategyResponseProviderTypeHTTP`    | http                                                                             |