# BillingInvoiceStandardStatus

Current lifecycle status of the invoice.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingInvoiceStandardStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.BillingInvoiceStandardStatus("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `BillingInvoiceStandardStatusDraft`             | draft                                           |
| `BillingInvoiceStandardStatusIssuing`           | issuing                                         |
| `BillingInvoiceStandardStatusIssued`            | issued                                          |
| `BillingInvoiceStandardStatusPaymentProcessing` | payment_processing                              |
| `BillingInvoiceStandardStatusOverdue`           | overdue                                         |
| `BillingInvoiceStandardStatusPaid`              | paid                                            |
| `BillingInvoiceStandardStatusUncollectible`     | uncollectible                                   |
| `BillingInvoiceStandardStatusVoided`            | voided                                          |