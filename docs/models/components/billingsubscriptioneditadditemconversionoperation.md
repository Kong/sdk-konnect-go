# BillingSubscriptionEditAddItemConversionOperation

The arithmetic operation to apply to the raw metered quantity.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionEditAddItemConversionOperationDivide

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionEditAddItemConversionOperation("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `BillingSubscriptionEditAddItemConversionOperationDivide`   | divide                                                      |
| `BillingSubscriptionEditAddItemConversionOperationMultiply` | multiply                                                    |