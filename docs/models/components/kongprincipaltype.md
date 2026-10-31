# KongPrincipalType

Classifies the kind of principal. `agent` is a deployed AI agent,
`user` is a human in the organization.
When omitted or null, the principal is not classified.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.KongPrincipalTypeAgent

// Open enum: custom values can be created with a direct type cast
custom := components.KongPrincipalType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `KongPrincipalTypeAgent` | agent                    |
| `KongPrincipalTypeUser`  | user                     |