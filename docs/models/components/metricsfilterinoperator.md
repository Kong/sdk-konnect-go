# MetricsFilterInOperator

The type of filter to apply.
  - `in` filters will limit results to only the specified values
  - `not_in` filters will exclude the specified values


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MetricsFilterInOperatorIn

// Open enum: custom values can be created with a direct type cast
custom := components.MetricsFilterInOperator("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `MetricsFilterInOperatorIn`    | in                             |
| `MetricsFilterInOperatorNotIn` | not_in                         |