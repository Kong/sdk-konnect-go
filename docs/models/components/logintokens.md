# LoginTokens

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.LoginTokensAccessToken

// Open enum: custom values can be created with a direct type cast
custom := components.LoginTokens("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `LoginTokensAccessToken`   | access_token               |
| `LoginTokensIDToken`       | id_token                   |
| `LoginTokensIntrospection` | introspection              |
| `LoginTokensRefreshToken`  | refresh_token              |
| `LoginTokensTokens`        | tokens                     |