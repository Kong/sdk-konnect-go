# BillingAppCapabilityType1

Type of the capability.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingAppCapabilityType1ReportUsage

// Open enum: custom values can be created with a direct type cast
custom := components.BillingAppCapabilityType1("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `BillingAppCapabilityType1ReportUsage`      | report_usage                                |
| `BillingAppCapabilityType1ReportEvents`     | report_events                               |
| `BillingAppCapabilityType1CalculateTax`     | calculate_tax                               |
| `BillingAppCapabilityType1InvoiceCustomers` | invoice_customers                           |
| `BillingAppCapabilityType1CollectPayments`  | collect_payments                            |