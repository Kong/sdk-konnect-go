# BillingInstalledAppExternalInvoicingDefinitionType

Type of the app.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingInstalledAppExternalInvoicingDefinitionTypeSandbox

// Open enum: custom values can be created with a direct type cast
custom := components.BillingInstalledAppExternalInvoicingDefinitionType("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `BillingInstalledAppExternalInvoicingDefinitionTypeSandbox`           | sandbox                                                               |
| `BillingInstalledAppExternalInvoicingDefinitionTypeStripe`            | stripe                                                                |
| `BillingInstalledAppExternalInvoicingDefinitionTypeExternalInvoicing` | external_invoicing                                                    |