# ConversionOperation

The arithmetic operation to apply to the raw metered quantity.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ConversionOperationDivide

// Open enum: custom values can be created with a direct type cast
custom := components.ConversionOperation("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `ConversionOperationDivide`   | divide                        |
| `ConversionOperationMultiply` | multiply                      |