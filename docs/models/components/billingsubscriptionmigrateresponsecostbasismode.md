# BillingSubscriptionMigrateResponseCostBasisMode

Controls whether custom-currency cost bases are resolved dynamically or pinned
when their currency pair is introduced to the subscription.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionMigrateResponseCostBasisModeDynamic

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionMigrateResponseCostBasisMode("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `BillingSubscriptionMigrateResponseCostBasisModeDynamic` | dynamic                                                  |
| `BillingSubscriptionMigrateResponseCostBasisModePinned`  | pinned                                                   |