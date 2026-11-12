# BillingSubscriptionChangeResponseMode

How pro-rating is calculated when enabled.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionChangeResponseModeNoProration

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionChangeResponseMode("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `BillingSubscriptionChangeResponseModeNoProration`   | no_proration                                         |
| `BillingSubscriptionChangeResponseModeProratePrices` | prorate_prices                                       |