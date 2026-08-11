# ClientAuth

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ClientAuthClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.ClientAuth("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `ClientAuthClientSecretBasic`       | client_secret_basic                 |
| `ClientAuthClientSecretPost`        | client_secret_post                  |
| `ClientAuthClientSecretJwt`         | client_secret_jwt                   |
| `ClientAuthPrivateKeyJwt`           | private_key_jwt                     |
| `ClientAuthTLSClientAuth`           | tls_client_auth                     |
| `ClientAuthSelfSignedTLSClientAuth` | self_signed_tls_client_auth         |
| `ClientAuthNone`                    | none                                |