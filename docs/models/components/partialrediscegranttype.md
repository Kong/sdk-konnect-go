# PartialRedisCeGrantType

OAuth 2.0 grant type used to request access tokens.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PartialRedisCeGrantTypeClientCredentials

// Open enum: custom values can be created with a direct type cast
custom := components.PartialRedisCeGrantType("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `PartialRedisCeGrantTypeClientCredentials` | client_credentials                         |
| `PartialRedisCeGrantTypePassword`          | password                                   |