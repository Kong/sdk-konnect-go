# BillingSubscriptionSubscriptionMode

How pro-rating is calculated when enabled.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionSubscriptionModeNoProration

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionSubscriptionMode("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `BillingSubscriptionSubscriptionModeNoProration`   | no_proration                                       |
| `BillingSubscriptionSubscriptionModeProratePrices` | prorate_prices                                     |