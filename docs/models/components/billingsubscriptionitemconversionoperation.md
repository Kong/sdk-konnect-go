# BillingSubscriptionItemConversionOperation

The arithmetic operation to apply to the raw metered quantity.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionItemConversionOperationDivide

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionItemConversionOperation("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `BillingSubscriptionItemConversionOperationDivide`   | divide                                               |
| `BillingSubscriptionItemConversionOperationMultiply` | multiply                                             |