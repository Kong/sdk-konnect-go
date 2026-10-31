# IdentityProviderType

Specifies the type of identity provider.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.IdentityProviderTypeOidc

// Open enum: custom values can be created with a direct type cast
custom := components.IdentityProviderType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `IdentityProviderTypeOidc` | oidc                       |
| `IdentityProviderTypeSaml` | saml                       |