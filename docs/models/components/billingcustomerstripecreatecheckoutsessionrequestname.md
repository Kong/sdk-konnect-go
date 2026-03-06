# BillingCustomerStripeCreateCheckoutSessionRequestName

Whether to save the customer name to customer.name.

Defaults to "never".

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingCustomerStripeCreateCheckoutSessionRequestNameAuto

// Open enum: custom values can be created with a direct type cast
custom := components.BillingCustomerStripeCreateCheckoutSessionRequestName("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `BillingCustomerStripeCreateCheckoutSessionRequestNameAuto`  | auto                                                         |
| `BillingCustomerStripeCreateCheckoutSessionRequestNameNever` | never                                                        |