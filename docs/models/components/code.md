# Code

Machine-readable error code.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CodeUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.Code("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `CodeUnknown`          | unknown                |
| `CodeCustomerNotFound` | customer_not_found     |