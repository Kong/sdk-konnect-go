# BillingUsageBasedChargeSettlementMode

Settlement mode of the charge.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingUsageBasedChargeSettlementModeCreditThenInvoice

// Open enum: custom values can be created with a direct type cast
custom := components.BillingUsageBasedChargeSettlementMode("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `BillingUsageBasedChargeSettlementModeCreditThenInvoice` | credit_then_invoice                                      |
| `BillingUsageBasedChargeSettlementModeCreditOnly`        | credit_only                                              |