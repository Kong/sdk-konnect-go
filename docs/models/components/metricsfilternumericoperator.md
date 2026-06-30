# MetricsFilterNumericOperator

The type of filter to apply to numeric field values.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MetricsFilterNumericOperatorEqual

// Open enum: custom values can be created with a direct type cast
custom := components.MetricsFilterNumericOperator("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `MetricsFilterNumericOperatorEqual`            | =                                              |
| `MetricsFilterNumericOperatorNotEqual`         | !=                                             |
| `MetricsFilterNumericOperatorLessThan`         | <                                              |
| `MetricsFilterNumericOperatorGreaterThan`      | >                                              |
| `MetricsFilterNumericOperatorLessThanEqual`    | <=                                             |
| `MetricsFilterNumericOperatorGreaterThanEqual` | >=                                             |