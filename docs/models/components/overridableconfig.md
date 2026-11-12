# OverridableConfig

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.OverridableConfigClientID

// Open enum: custom values can be created with a direct type cast
custom := components.OverridableConfig("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `OverridableConfigClientID`              | client_id                                |
| `OverridableConfigClientSecret`          | client_secret                            |
| `OverridableConfigAuthorizationEndpoint` | authorization_endpoint                   |
| `OverridableConfigTokenEndpoint`         | token_endpoint                           |