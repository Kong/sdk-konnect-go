# PersonalAccessTokenState

State of the personal access token.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PersonalAccessTokenStateActive

// Open enum: custom values can be created with a direct type cast
custom := components.PersonalAccessTokenState("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PersonalAccessTokenStateActive`  | ACTIVE                            |
| `PersonalAccessTokenStateRevoked` | REVOKED                           |
| `PersonalAccessTokenStateExpired` | EXPIRED                           |