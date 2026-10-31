# RequestsFilterType

The type of filter to apply.  `in` filters will limit results to only the specified values, while `not_in` filters will exclude the specified values.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.RequestsFilterTypeIn

// Open enum: custom values can be created with a direct type cast
custom := components.RequestsFilterType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `RequestsFilterTypeIn`    | in                        |
| `RequestsFilterTypeNotIn` | not_in                    |