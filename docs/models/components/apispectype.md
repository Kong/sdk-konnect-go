# APISpecType

The type of specification being stored. This allows us to render the specification correctly.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.APISpecTypeOas2

// Open enum: custom values can be created with a direct type cast
custom := components.APISpecType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `APISpecTypeOas2`     | oas2                  |
| `APISpecTypeOas3`     | oas3                  |
| `APISpecTypeAsyncapi` | asyncapi              |