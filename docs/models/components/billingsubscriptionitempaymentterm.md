# BillingSubscriptionItemPaymentTerm

The payment term of the rate card. In advance payment term can only be used for
flat prices.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionItemPaymentTermInAdvance

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionItemPaymentTerm("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `BillingSubscriptionItemPaymentTermInAdvance` | in_advance                                    |
| `BillingSubscriptionItemPaymentTermInArrears` | in_arrears                                    |