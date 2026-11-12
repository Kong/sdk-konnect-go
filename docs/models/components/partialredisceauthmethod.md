# PartialRedisCeAuthMethod

Client authentication method used against the token endpoint.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PartialRedisCeAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.PartialRedisCeAuthMethod("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `PartialRedisCeAuthMethodClientSecretBasic` | client_secret_basic                         |
| `PartialRedisCeAuthMethodClientSecretJwt`   | client_secret_jwt                           |
| `PartialRedisCeAuthMethodClientSecretPost`  | client_secret_post                          |