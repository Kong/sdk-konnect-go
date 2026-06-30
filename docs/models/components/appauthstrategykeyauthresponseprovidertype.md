# AppAuthStrategyKeyAuthResponseProviderType

The type of DCR provider.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AppAuthStrategyKeyAuthResponseProviderTypeAuth0

// Open enum: custom values can be created with a direct type cast
custom := components.AppAuthStrategyKeyAuthResponseProviderType("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `AppAuthStrategyKeyAuthResponseProviderTypeAuth0`   | auth0                                               |
| `AppAuthStrategyKeyAuthResponseProviderTypeAzureAd` | azureAd                                             |
| `AppAuthStrategyKeyAuthResponseProviderTypeCurity`  | curity                                              |
| `AppAuthStrategyKeyAuthResponseProviderTypeOkta`    | okta                                                |
| `AppAuthStrategyKeyAuthResponseProviderTypeHTTP`    | http                                                |