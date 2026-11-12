# PartialRedisEeAuthMethod

Client authentication method used against the token endpoint.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PartialRedisEeAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.PartialRedisEeAuthMethod("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `PartialRedisEeAuthMethodClientSecretBasic` | client_secret_basic                         |
| `PartialRedisEeAuthMethodClientSecretJwt`   | client_secret_jwt                           |
| `PartialRedisEeAuthMethodClientSecretPost`  | client_secret_post                          |