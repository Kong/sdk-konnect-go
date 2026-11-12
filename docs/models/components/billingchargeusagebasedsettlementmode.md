# BillingChargeUsageBasedSettlementMode

Settlement mode of the charge.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingChargeUsageBasedSettlementModeCreditThenInvoice

// Open enum: custom values can be created with a direct type cast
custom := components.BillingChargeUsageBasedSettlementMode("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `BillingChargeUsageBasedSettlementModeCreditThenInvoice` | credit_then_invoice                                      |
| `BillingChargeUsageBasedSettlementModeCreditOnly`        | credit_only                                              |