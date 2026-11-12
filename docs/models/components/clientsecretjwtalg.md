# ClientSecretJwtAlg

Signing algorithm used for `client_secret_jwt` client authentication.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ClientSecretJwtAlgHs256

// Open enum: custom values can be created with a direct type cast
custom := components.ClientSecretJwtAlg("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ClientSecretJwtAlgHs256` | HS256                     |
| `ClientSecretJwtAlgHs512` | HS512                     |