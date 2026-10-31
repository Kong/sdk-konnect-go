# IgnoreSignature

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.IgnoreSignatureAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.IgnoreSignature("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `IgnoreSignatureAuthorizationCode` | authorization_code                 |
| `IgnoreSignatureClientCredentials` | client_credentials                 |
| `IgnoreSignatureIntrospection`     | introspection                      |
| `IgnoreSignaturePassword`          | password                           |
| `IgnoreSignatureRefreshToken`      | refresh_token                      |
| `IgnoreSignatureSession`           | session                            |
| `IgnoreSignatureUserinfo`          | userinfo                           |