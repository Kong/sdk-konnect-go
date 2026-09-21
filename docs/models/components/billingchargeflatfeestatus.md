# BillingChargeFlatFeeStatus

The lifecycle status of the charge.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingChargeFlatFeeStatusCreated

// Open enum: custom values can be created with a direct type cast
custom := components.BillingChargeFlatFeeStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `BillingChargeFlatFeeStatusCreated` | created                             |
| `BillingChargeFlatFeeStatusActive`  | active                              |
| `BillingChargeFlatFeeStatusFinal`   | final                               |
| `BillingChargeFlatFeeStatusDeleted` | deleted                             |