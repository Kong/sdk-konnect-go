# SubscriptionAddonRateCardPaymentTerm

The payment term of the rate card. In advance payment term can only be used for
flat prices.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SubscriptionAddonRateCardPaymentTermInAdvance

// Open enum: custom values can be created with a direct type cast
custom := components.SubscriptionAddonRateCardPaymentTerm("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `SubscriptionAddonRateCardPaymentTermInAdvance` | in_advance                                      |
| `SubscriptionAddonRateCardPaymentTermInArrears` | in_arrears                                      |