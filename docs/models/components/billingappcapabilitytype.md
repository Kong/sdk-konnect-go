# BillingAppCapabilityType

Type of the capability.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingAppCapabilityTypeReportUsage

// Open enum: custom values can be created with a direct type cast
custom := components.BillingAppCapabilityType("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `BillingAppCapabilityTypeReportUsage`      | report_usage                               |
| `BillingAppCapabilityTypeReportEvents`     | report_events                              |
| `BillingAppCapabilityTypeCalculateTax`     | calculate_tax                              |
| `BillingAppCapabilityTypeInvoiceCustomers` | invoice_customers                          |
| `BillingAppCapabilityTypeCollectPayments`  | collect_payments                           |