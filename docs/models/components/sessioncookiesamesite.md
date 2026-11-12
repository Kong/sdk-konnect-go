# SessionCookieSameSite

Controls whether a cookie is sent with cross-origin requests, providing some protection against cross-site request forgery attacks.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SessionCookieSameSiteDefault

// Open enum: custom values can be created with a direct type cast
custom := components.SessionCookieSameSite("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `SessionCookieSameSiteDefault` | Default                        |
| `SessionCookieSameSiteLax`     | Lax                            |
| `SessionCookieSameSiteNone`    | None                           |
| `SessionCookieSameSiteStrict`  | Strict                         |