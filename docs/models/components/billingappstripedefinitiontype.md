# BillingAppStripeDefinitionType

Type of the app.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingAppStripeDefinitionTypeSandbox

// Open enum: custom values can be created with a direct type cast
custom := components.BillingAppStripeDefinitionType("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `BillingAppStripeDefinitionTypeSandbox`           | sandbox                                           |
| `BillingAppStripeDefinitionTypeStripe`            | stripe                                            |
| `BillingAppStripeDefinitionTypeExternalInvoicing` | external_invoicing                                |