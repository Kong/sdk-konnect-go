# PartialVectordbAuthMethod

Client authentication method used against the token endpoint.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PartialVectordbAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.PartialVectordbAuthMethod("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `PartialVectordbAuthMethodClientSecretBasic` | client_secret_basic                          |
| `PartialVectordbAuthMethodClientSecretJwt`   | client_secret_jwt                            |
| `PartialVectordbAuthMethodClientSecretPost`  | client_secret_post                           |