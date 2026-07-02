# MetricsFilterEmptyOperator

The type of filter to apply.
Empty values are `null` or undefined.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MetricsFilterEmptyOperatorEmpty

// Open enum: custom values can be created with a direct type cast
custom := components.MetricsFilterEmptyOperator("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `MetricsFilterEmptyOperatorEmpty`    | empty                                |
| `MetricsFilterEmptyOperatorNotEmpty` | not_empty                            |