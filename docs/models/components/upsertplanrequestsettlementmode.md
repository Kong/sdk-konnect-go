# UpsertPlanRequestSettlementMode

Settlement mode for the plan. When omitted, the existing settlement mode is
preserved.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.UpsertPlanRequestSettlementModeCreditThenInvoice

// Open enum: custom values can be created with a direct type cast
custom := components.UpsertPlanRequestSettlementMode("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `UpsertPlanRequestSettlementModeCreditThenInvoice` | credit_then_invoice                                |
| `UpsertPlanRequestSettlementModeCreditOnly`        | credit_only                                        |