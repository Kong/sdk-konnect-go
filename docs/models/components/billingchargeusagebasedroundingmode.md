# BillingChargeUsageBasedRoundingMode

The rounding mode applied to the converted quantity for invoicing.

Defaults to none (no rounding). Entitlement checks always use the precise
(unrounded) value.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingChargeUsageBasedRoundingModeCeiling

// Open enum: custom values can be created with a direct type cast
custom := components.BillingChargeUsageBasedRoundingMode("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `BillingChargeUsageBasedRoundingModeCeiling` | ceiling                                      |
| `BillingChargeUsageBasedRoundingModeFloor`   | floor                                        |
| `BillingChargeUsageBasedRoundingModeHalfUp`  | half_up                                      |
| `BillingChargeUsageBasedRoundingModeNone`    | none                                         |