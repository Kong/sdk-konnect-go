# BillingSubscriptionItemRoundingMode

The rounding mode applied to the converted quantity for invoicing.

Defaults to none (no rounding). Entitlement checks always use the precise
(unrounded) value.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionItemRoundingModeCeiling

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionItemRoundingMode("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `BillingSubscriptionItemRoundingModeCeiling` | ceiling                                      |
| `BillingSubscriptionItemRoundingModeFloor`   | floor                                        |
| `BillingSubscriptionItemRoundingModeHalfUp`  | half_up                                      |
| `BillingSubscriptionItemRoundingModeNone`    | none                                         |