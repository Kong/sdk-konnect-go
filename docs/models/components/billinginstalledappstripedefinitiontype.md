# BillingInstalledAppStripeDefinitionType

Type of the app.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingInstalledAppStripeDefinitionTypeSandbox

// Open enum: custom values can be created with a direct type cast
custom := components.BillingInstalledAppStripeDefinitionType("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `BillingInstalledAppStripeDefinitionTypeSandbox`           | sandbox                                                    |
| `BillingInstalledAppStripeDefinitionTypeStripe`            | stripe                                                     |
| `BillingInstalledAppStripeDefinitionTypeExternalInvoicing` | external_invoicing                                         |