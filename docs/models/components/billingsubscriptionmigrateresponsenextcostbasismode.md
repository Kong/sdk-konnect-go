# BillingSubscriptionMigrateResponseNextCostBasisMode

Controls whether custom-currency cost bases are resolved dynamically or pinned
when their currency pair is introduced to the subscription.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionMigrateResponseNextCostBasisModeDynamic

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionMigrateResponseNextCostBasisMode("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `BillingSubscriptionMigrateResponseNextCostBasisModeDynamic` | dynamic                                                      |
| `BillingSubscriptionMigrateResponseNextCostBasisModePinned`  | pinned                                                       |