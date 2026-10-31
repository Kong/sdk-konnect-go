# LoginAction

What to do after successful login: - `upstream`: proxy request to upstream service - `response`: terminate request with a response - `redirect`: redirect to a different location.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.LoginActionRedirect

// Open enum: custom values can be created with a direct type cast
custom := components.LoginAction("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `LoginActionRedirect` | redirect              |
| `LoginActionResponse` | response              |
| `LoginActionUpstream` | upstream              |