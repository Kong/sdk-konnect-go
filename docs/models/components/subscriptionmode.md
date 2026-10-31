# SubscriptionMode

How pro-rating is calculated when enabled.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SubscriptionModeNoProration

// Open enum: custom values can be created with a direct type cast
custom := components.SubscriptionMode("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `SubscriptionModeNoProration`   | no_proration                    |
| `SubscriptionModeProratePrices` | prorate_prices                  |