# Shipping

Whether to save shipping information to customer.shipping.

Defaults to "never".

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.ShippingAuto

// Open enum: custom values can be created with a direct type cast
custom := metering.Shipping("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `ShippingAuto`  | auto            |
| `ShippingNever` | never           |