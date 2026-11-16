# AuthMethod

Client authentication method used against the token endpoint.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AuthMethod("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `AuthMethodClientSecretBasic` | client_secret_basic           |
| `AuthMethodClientSecretJwt`   | client_secret_jwt             |
| `AuthMethodClientSecretPost`  | client_secret_post            |