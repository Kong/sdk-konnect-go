# APISpecAPISpecType

The type of specification being stored. This allows us to render the specification correctly.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.APISpecAPISpecTypeOas2

// Open enum: custom values can be created with a direct type cast
custom := components.APISpecAPISpecType("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `APISpecAPISpecTypeOas2`     | oas2                         |
| `APISpecAPISpecTypeOas3`     | oas3                         |
| `APISpecAPISpecTypeAsyncapi` | asyncapi                     |