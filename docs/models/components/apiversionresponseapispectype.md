# APIVersionResponseAPISpecType

The type of specification being stored. This allows us to render the specification correctly.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.APIVersionResponseAPISpecTypeOas2

// Open enum: custom values can be created with a direct type cast
custom := components.APIVersionResponseAPISpecType("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `APIVersionResponseAPISpecTypeOas2`     | oas2                                    |
| `APIVersionResponseAPISpecTypeOas3`     | oas3                                    |
| `APIVersionResponseAPISpecTypeAsyncapi` | asyncapi                                |