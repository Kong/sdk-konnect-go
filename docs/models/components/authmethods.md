# AuthMethods

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AuthMethodsAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AuthMethods("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `AuthMethodsAuthorizationCode` | authorization_code             |
| `AuthMethodsBearer`            | bearer                         |
| `AuthMethodsClientCredentials` | client_credentials             |
| `AuthMethodsIntrospection`     | introspection                  |
| `AuthMethodsKongOauth2`        | kong_oauth2                    |
| `AuthMethodsPassword`          | password                       |
| `AuthMethodsRefreshToken`      | refresh_token                  |
| `AuthMethodsSession`           | session                        |
| `AuthMethodsUserinfo`          | userinfo                       |