# BillingPlanSettlementMode

Settlement mode for plan.

Values:

- `credit_then_invoice`: Credits are applied first, then any remainder is
invoiced.
- `credit_only`: Usage is settled exclusively against credits.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingPlanSettlementModeCreditThenInvoice

// Open enum: custom values can be created with a direct type cast
custom := components.BillingPlanSettlementMode("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `BillingPlanSettlementModeCreditThenInvoice` | credit_then_invoice                          |
| `BillingPlanSettlementModeCreditOnly`        | credit_only                                  |