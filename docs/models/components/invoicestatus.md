# InvoiceStatus

Current lifecycle status of the invoice.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.InvoiceStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.InvoiceStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `InvoiceStatusDraft`             | draft                            |
| `InvoiceStatusIssuing`           | issuing                          |
| `InvoiceStatusIssued`            | issued                           |
| `InvoiceStatusPaymentProcessing` | payment_processing               |
| `InvoiceStatusOverdue`           | overdue                          |
| `InvoiceStatusPaid`              | paid                             |
| `InvoiceStatusUncollectible`     | uncollectible                    |
| `InvoiceStatusVoided`            | voided                           |