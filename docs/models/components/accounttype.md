# AccountType

The authorization strategy used to register the user.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AccountTypeGoogle

// Open enum: custom values can be created with a direct type cast
custom := components.AccountType("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `AccountTypeGoogle`    | google                 |
| `AccountTypeGithub`    | github                 |
| `AccountTypeMicrosoft` | microsoft              |
| `AccountTypeBasicAuth` | basic-auth             |
| `AccountTypeSso`       | sso                    |