# In

The location of the parameter in the request.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.InQuery

// Open enum: custom values can be created with a direct type cast
custom := components.In("custom_value")
```


## Values

| Name       | Value      |
| ---------- | ---------- |
| `InQuery`  | query      |
| `InPath`   | path       |
| `InHeader` | header     |
| `InBody`   | body       |