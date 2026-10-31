# CredentialProviderTemplateBearerMethod

Method used to present the bearer token to the provider.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CredentialProviderTemplateBearerMethodHeader

// Open enum: custom values can be created with a direct type cast
custom := components.CredentialProviderTemplateBearerMethod("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `CredentialProviderTemplateBearerMethodHeader` | header                                         |
| `CredentialProviderTemplateBearerMethodForm`   | form                                           |