# CreateChargeUsageBasedRequestSettlementMode

Settlement mode of the charge.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateChargeUsageBasedRequestSettlementModeCreditThenInvoice

// Open enum: custom values can be created with a direct type cast
custom := components.CreateChargeUsageBasedRequestSettlementMode("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `CreateChargeUsageBasedRequestSettlementModeCreditThenInvoice` | credit_then_invoice                                            |
| `CreateChargeUsageBasedRequestSettlementModeCreditOnly`        | credit_only                                                    |