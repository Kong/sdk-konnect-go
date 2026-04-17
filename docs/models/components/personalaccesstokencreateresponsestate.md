# PersonalAccessTokenCreateResponseState

State of the personal access token.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PersonalAccessTokenCreateResponseStateActive

// Open enum: custom values can be created with a direct type cast
custom := components.PersonalAccessTokenCreateResponseState("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `PersonalAccessTokenCreateResponseStateActive`  | ACTIVE                                          |
| `PersonalAccessTokenCreateResponseStateRevoked` | REVOKED                                         |
| `PersonalAccessTokenCreateResponseStateExpired` | EXPIRED                                         |