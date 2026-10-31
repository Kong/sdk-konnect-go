# BillingSubscriptionMode

How pro-rating is calculated when enabled.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionModeNoProration

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionMode("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `BillingSubscriptionModeNoProration`   | no_proration                           |
| `BillingSubscriptionModeProratePrices` | prorate_prices                         |