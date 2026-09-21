# BillingInvoiceLineUsageDiscountReason

The reason this discount was applied.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingInvoiceLineUsageDiscountReasonMaximumSpend

// Open enum: custom values can be created with a direct type cast
custom := components.BillingInvoiceLineUsageDiscountReason("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `BillingInvoiceLineUsageDiscountReasonMaximumSpend`       | maximum_spend                                             |
| `BillingInvoiceLineUsageDiscountReasonRatecardPercentage` | ratecard_percentage                                       |
| `BillingInvoiceLineUsageDiscountReasonRatecardUsage`      | ratecard_usage                                            |