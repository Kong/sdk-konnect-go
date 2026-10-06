# TokenHeadersGrants

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.TokenHeadersGrantsAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.TokenHeadersGrants("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `TokenHeadersGrantsAuthorizationCode` | authorization_code                    |
| `TokenHeadersGrantsClientCredentials` | client_credentials                    |
| `TokenHeadersGrantsPassword`          | password                              |
| `TokenHeadersGrantsRefreshToken`      | refresh_token                         |