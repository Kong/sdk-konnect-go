# AppAuthStrategyKeyAuthResponseAppAuthStrategyProviderType

The type of DCR provider.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AppAuthStrategyKeyAuthResponseAppAuthStrategyProviderTypeAuth0

// Open enum: custom values can be created with a direct type cast
custom := components.AppAuthStrategyKeyAuthResponseAppAuthStrategyProviderType("custom_value")
```


## Values

| Name                                                                    | Value                                                                   |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `AppAuthStrategyKeyAuthResponseAppAuthStrategyProviderTypeAuth0`        | auth0                                                                   |
| `AppAuthStrategyKeyAuthResponseAppAuthStrategyProviderTypeAzureAd`      | azureAd                                                                 |
| `AppAuthStrategyKeyAuthResponseAppAuthStrategyProviderTypeCurity`       | curity                                                                  |
| `AppAuthStrategyKeyAuthResponseAppAuthStrategyProviderTypeOkta`         | okta                                                                    |
| `AppAuthStrategyKeyAuthResponseAppAuthStrategyProviderTypeHTTP`         | http                                                                    |
| `AppAuthStrategyKeyAuthResponseAppAuthStrategyProviderTypeKongIdentity` | kongIdentity                                                            |