# CreateChargeFlatFeeRequestSettlementMode

Settlement mode of the charge.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateChargeFlatFeeRequestSettlementModeCreditThenInvoice

// Open enum: custom values can be created with a direct type cast
custom := components.CreateChargeFlatFeeRequestSettlementMode("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `CreateChargeFlatFeeRequestSettlementModeCreditThenInvoice` | credit_then_invoice                                         |
| `CreateChargeFlatFeeRequestSettlementModeCreditOnly`        | credit_only                                                 |