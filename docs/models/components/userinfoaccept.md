# UserinfoAccept

The value of `Accept` header for user info requests: - `application/json`: user info response as JSON - `application/jwt`: user info response as JWT (from the obsolete IETF draft document).

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.UserinfoAcceptApplicationJSON

// Open enum: custom values can be created with a direct type cast
custom := components.UserinfoAccept("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `UserinfoAcceptApplicationJSON` | application/json                |
| `UserinfoAcceptApplicationJwt`  | application/jwt                 |