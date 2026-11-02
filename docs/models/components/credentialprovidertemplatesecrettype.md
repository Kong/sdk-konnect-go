# CredentialProviderTemplateSecretType

The OAuth grant type used to obtain credentials for this provider template.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CredentialProviderTemplateSecretTypeAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.CredentialProviderTemplateSecretType("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `CredentialProviderTemplateSecretTypeAuthorizationCode` | authorization_code                                      |
| `CredentialProviderTemplateSecretTypeStaticSecret`      | static_secret                                           |