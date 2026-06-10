# Promotions

Enables collection of promotional communication consent.

Only available to US merchants. When set to "auto", Checkout determines whether
to show the option based on the customer's locale.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.PromotionsAuto

// Open enum: custom values can be created with a direct type cast
custom := metering.Promotions("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `PromotionsAuto` | auto             |
| `PromotionsNone` | none             |