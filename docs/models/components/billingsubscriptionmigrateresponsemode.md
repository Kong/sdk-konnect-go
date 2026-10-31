# BillingSubscriptionMigrateResponseMode

How pro-rating is calculated when enabled.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionMigrateResponseModeNoProration

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionMigrateResponseMode("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `BillingSubscriptionMigrateResponseModeNoProration`   | no_proration                                          |
| `BillingSubscriptionMigrateResponseModeProratePrices` | prorate_prices                                        |