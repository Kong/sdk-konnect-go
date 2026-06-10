# BillingAppStripeType

Type of the app.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.BillingAppStripeTypeSandbox

// Open enum: custom values can be created with a direct type cast
custom := metering.BillingAppStripeType("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `BillingAppStripeTypeSandbox`           | sandbox                                 |
| `BillingAppStripeTypeStripe`            | stripe                                  |
| `BillingAppStripeTypeExternalInvoicing` | external_invoicing                      |