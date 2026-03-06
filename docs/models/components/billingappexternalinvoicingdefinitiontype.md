# BillingAppExternalInvoicingDefinitionType

Type of the app.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingAppExternalInvoicingDefinitionTypeSandbox

// Open enum: custom values can be created with a direct type cast
custom := components.BillingAppExternalInvoicingDefinitionType("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `BillingAppExternalInvoicingDefinitionTypeSandbox`           | sandbox                                                      |
| `BillingAppExternalInvoicingDefinitionTypeStripe`            | stripe                                                       |
| `BillingAppExternalInvoicingDefinitionTypeExternalInvoicing` | external_invoicing                                           |