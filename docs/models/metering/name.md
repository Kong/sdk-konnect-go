# Name

Whether to save the customer name to customer.name.

Defaults to "never".

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.NameAuto

// Open enum: custom values can be created with a direct type cast
custom := metering.Name("custom_value")
```


## Values

| Name        | Value       |
| ----------- | ----------- |
| `NameAuto`  | auto        |
| `NameNever` | never       |