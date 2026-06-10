# PaymentTerm

The payment term of the rate card. In advance payment term can only be used for
flat prices.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.PaymentTermInAdvance

// Open enum: custom values can be created with a direct type cast
custom := metering.PaymentTerm("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `PaymentTermInAdvance` | in_advance             |
| `PaymentTermInArrears` | in_arrears             |