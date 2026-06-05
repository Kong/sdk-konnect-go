# BillingUsageBasedChargeStatus

The lifecycle status of the charge.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingUsageBasedChargeStatusCreated

// Open enum: custom values can be created with a direct type cast
custom := components.BillingUsageBasedChargeStatus("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `BillingUsageBasedChargeStatusCreated` | created                                |
| `BillingUsageBasedChargeStatusActive`  | active                                 |
| `BillingUsageBasedChargeStatusFinal`   | final                                  |
| `BillingUsageBasedChargeStatusDeleted` | deleted                                |