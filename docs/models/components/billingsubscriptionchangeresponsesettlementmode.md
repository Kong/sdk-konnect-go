# BillingSubscriptionChangeResponseSettlementMode

Settlement mode for billing.

Values:

- `credit_then_invoice`: Credits are applied first, then any remainder is
invoiced.
- `credit_only`: Usage is settled exclusively against credits.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionChangeResponseSettlementModeCreditThenInvoice

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionChangeResponseSettlementMode("custom_value")
```


## Values

| Name                                                               | Value                                                              |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `BillingSubscriptionChangeResponseSettlementModeCreditThenInvoice` | credit_then_invoice                                                |
| `BillingSubscriptionChangeResponseSettlementModeCreditOnly`        | credit_only                                                        |