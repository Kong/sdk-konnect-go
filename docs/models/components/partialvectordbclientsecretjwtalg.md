# PartialVectordbClientSecretJwtAlg

Signing algorithm used for `client_secret_jwt` client authentication.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PartialVectordbClientSecretJwtAlgHs256

// Open enum: custom values can be created with a direct type cast
custom := components.PartialVectordbClientSecretJwtAlg("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `PartialVectordbClientSecretJwtAlgHs256` | HS256                                    |
| `PartialVectordbClientSecretJwtAlgHs512` | HS512                                    |