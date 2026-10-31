# BillingSubscriptionMigrateResponseNextMode

How pro-rating is calculated when enabled.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionMigrateResponseNextModeNoProration

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionMigrateResponseNextMode("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `BillingSubscriptionMigrateResponseNextModeNoProration`   | no_proration                                              |
| `BillingSubscriptionMigrateResponseNextModeProratePrices` | prorate_prices                                            |