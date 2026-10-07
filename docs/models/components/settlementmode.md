# SettlementMode

Settlement mode of the charge.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SettlementModeCreditThenInvoice

// Open enum: custom values can be created with a direct type cast
custom := components.SettlementMode("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `SettlementModeCreditThenInvoice` | credit_then_invoice               |
| `SettlementModeCreditOnly`        | credit_only                       |