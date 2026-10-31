# AuthServerGrantType

OAuth 2.0 grant type

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AuthServerGrantTypeImplicit

// Open enum: custom values can be created with a direct type cast
custom := components.AuthServerGrantType("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `AuthServerGrantTypeImplicit`                             | implicit                                                  |
| `AuthServerGrantTypeClientCredentials`                    | client_credentials                                        |
| `AuthServerGrantTypeUrnIetfParamsOauthGrantTypeJwtBearer` | urn:ietf:params:oauth:grant-type:jwt-bearer               |