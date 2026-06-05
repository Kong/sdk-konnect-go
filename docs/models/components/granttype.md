# GrantType

The OAuth 2.0 grant type used for authorization (e.g., `authorization_code`).
Determines the flow the integration uses to request access tokens.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.GrantTypeAuthorizationCode
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `GrantTypeAuthorizationCode` | authorization_code           |