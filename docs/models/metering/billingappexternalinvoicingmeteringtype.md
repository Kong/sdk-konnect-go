# BillingAppExternalInvoicingMeteringType

Type of the app.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.BillingAppExternalInvoicingMeteringTypeSandbox

// Open enum: custom values can be created with a direct type cast
custom := metering.BillingAppExternalInvoicingMeteringType("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `BillingAppExternalInvoicingMeteringTypeSandbox`           | sandbox                                                    |
| `BillingAppExternalInvoicingMeteringTypeStripe`            | stripe                                                     |
| `BillingAppExternalInvoicingMeteringTypeExternalInvoicing` | external_invoicing                                         |