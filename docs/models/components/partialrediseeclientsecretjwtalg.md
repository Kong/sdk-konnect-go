# PartialRedisEeClientSecretJwtAlg

Signing algorithm used for `client_secret_jwt` client authentication.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PartialRedisEeClientSecretJwtAlgHs256

// Open enum: custom values can be created with a direct type cast
custom := components.PartialRedisEeClientSecretJwtAlg("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `PartialRedisEeClientSecretJwtAlgHs256` | HS256                                   |
| `PartialRedisEeClientSecretJwtAlgHs512` | HS512                                   |