# LoginMethods

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.LoginMethodsAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.LoginMethods("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `LoginMethodsAuthorizationCode` | authorization_code              |
| `LoginMethodsBearer`            | bearer                          |
| `LoginMethodsClientCredentials` | client_credentials              |
| `LoginMethodsIntrospection`     | introspection                   |
| `LoginMethodsKongOauth2`        | kong_oauth2                     |
| `LoginMethodsPassword`          | password                        |
| `LoginMethodsRefreshToken`      | refresh_token                   |
| `LoginMethodsSession`           | session                         |
| `LoginMethodsUserinfo`          | userinfo                        |