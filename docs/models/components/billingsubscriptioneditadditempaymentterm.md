# BillingSubscriptionEditAddItemPaymentTerm

The payment term of the rate card. In advance payment term can only be used for
flat prices.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionEditAddItemPaymentTermInAdvance

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionEditAddItemPaymentTerm("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `BillingSubscriptionEditAddItemPaymentTermInAdvance` | in_advance                                           |
| `BillingSubscriptionEditAddItemPaymentTermInArrears` | in_arrears                                           |