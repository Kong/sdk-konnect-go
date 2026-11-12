# BillingSubscriptionMigrateResponseSettlementMode

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

value := components.BillingSubscriptionMigrateResponseSettlementModeCreditThenInvoice

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionMigrateResponseSettlementMode("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `BillingSubscriptionMigrateResponseSettlementModeCreditThenInvoice` | credit_then_invoice                                                 |
| `BillingSubscriptionMigrateResponseSettlementModeCreditOnly`        | credit_only                                                         |