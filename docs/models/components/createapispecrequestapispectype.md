# CreateAPISpecRequestAPISpecType

The type of specification being stored. This allows us to render the specification correctly.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateAPISpecRequestAPISpecTypeOas2

// Open enum: custom values can be created with a direct type cast
custom := components.CreateAPISpecRequestAPISpecType("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `CreateAPISpecRequestAPISpecTypeOas2`     | oas2                                      |
| `CreateAPISpecRequestAPISpecTypeOas3`     | oas3                                      |
| `CreateAPISpecRequestAPISpecTypeAsyncapi` | asyncapi                                  |