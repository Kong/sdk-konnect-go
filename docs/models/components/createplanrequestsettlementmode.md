# CreatePlanRequestSettlementMode

Settlement mode for the plan. Defaults to `credit_then_invoice` when omitted.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreatePlanRequestSettlementModeCreditThenInvoice

// Open enum: custom values can be created with a direct type cast
custom := components.CreatePlanRequestSettlementMode("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `CreatePlanRequestSettlementModeCreditThenInvoice` | credit_then_invoice                                |
| `CreatePlanRequestSettlementModeCreditOnly`        | credit_only                                        |