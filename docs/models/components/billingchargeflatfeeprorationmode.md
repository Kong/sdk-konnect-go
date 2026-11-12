# BillingChargeFlatFeeProrationMode

The proration mode of the rate card.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingChargeFlatFeeProrationModeNoProration

// Open enum: custom values can be created with a direct type cast
custom := components.BillingChargeFlatFeeProrationMode("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `BillingChargeFlatFeeProrationModeNoProration`   | no_proration                                     |
| `BillingChargeFlatFeeProrationModeProratePrices` | prorate_prices                                   |