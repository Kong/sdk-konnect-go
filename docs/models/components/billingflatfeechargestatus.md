# BillingFlatFeeChargeStatus

The lifecycle status of the charge.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingFlatFeeChargeStatusCreated

// Open enum: custom values can be created with a direct type cast
custom := components.BillingFlatFeeChargeStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `BillingFlatFeeChargeStatusCreated` | created                             |
| `BillingFlatFeeChargeStatusActive`  | active                              |
| `BillingFlatFeeChargeStatusFinal`   | final                               |
| `BillingFlatFeeChargeStatusDeleted` | deleted                             |