# BillingChargeUsageBasedStatus

The lifecycle status of the charge.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingChargeUsageBasedStatusCreated

// Open enum: custom values can be created with a direct type cast
custom := components.BillingChargeUsageBasedStatus("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `BillingChargeUsageBasedStatusCreated` | created                                |
| `BillingChargeUsageBasedStatusActive`  | active                                 |
| `BillingChargeUsageBasedStatusFinal`   | final                                  |
| `BillingChargeUsageBasedStatusDeleted` | deleted                                |