# ProviderType

The type of DCR provider.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ProviderTypeAuth0

// Open enum: custom values can be created with a direct type cast
custom := components.ProviderType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ProviderTypeAuth0`        | auth0                      |
| `ProviderTypeAzureAd`      | azureAd                    |
| `ProviderTypeCurity`       | curity                     |
| `ProviderTypeOkta`         | okta                       |
| `ProviderTypeHTTP`         | http                       |
| `ProviderTypeKongIdentity` | kongIdentity               |