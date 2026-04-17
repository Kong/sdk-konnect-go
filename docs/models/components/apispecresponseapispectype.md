# APISpecResponseAPISpecType

The type of specification being stored. This allows us to render the specification correctly.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.APISpecResponseAPISpecTypeOas2

// Open enum: custom values can be created with a direct type cast
custom := components.APISpecResponseAPISpecType("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `APISpecResponseAPISpecTypeOas2`     | oas2                                 |
| `APISpecResponseAPISpecTypeOas3`     | oas3                                 |
| `APISpecResponseAPISpecTypeAsyncapi` | asyncapi                             |