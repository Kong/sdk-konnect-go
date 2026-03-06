# BillingAppStripeType

Type of the app.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingAppStripeTypeSandbox

// Open enum: custom values can be created with a direct type cast
custom := components.BillingAppStripeType("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `BillingAppStripeTypeSandbox`           | sandbox                                 |
| `BillingAppStripeTypeStripe`            | stripe                                  |
| `BillingAppStripeTypeExternalInvoicing` | external_invoicing                      |