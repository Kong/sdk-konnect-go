# PartialRedisEeGrantType

OAuth 2.0 grant type used to request access tokens.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PartialRedisEeGrantTypeClientCredentials

// Open enum: custom values can be created with a direct type cast
custom := components.PartialRedisEeGrantType("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `PartialRedisEeGrantTypeClientCredentials` | client_credentials                         |
| `PartialRedisEeGrantTypePassword`          | password                                   |