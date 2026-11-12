# DisableSession

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DisableSessionAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.DisableSession("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `DisableSessionAuthorizationCode` | authorization_code                |
| `DisableSessionBearer`            | bearer                            |
| `DisableSessionClientCredentials` | client_credentials                |
| `DisableSessionIntrospection`     | introspection                     |
| `DisableSessionKongOauth2`        | kong_oauth2                       |
| `DisableSessionPassword`          | password                          |
| `DisableSessionRefreshToken`      | refresh_token                     |
| `DisableSessionSession`           | session                           |
| `DisableSessionUserinfo`          | userinfo                          |