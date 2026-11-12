# BillingSubscriptionChangeResponseNextMode

How pro-rating is calculated when enabled.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionChangeResponseNextModeNoProration

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionChangeResponseNextMode("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `BillingSubscriptionChangeResponseNextModeNoProration`   | no_proration                                             |
| `BillingSubscriptionChangeResponseNextModeProratePrices` | prorate_prices                                           |