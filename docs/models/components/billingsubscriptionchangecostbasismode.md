# BillingSubscriptionChangeCostBasisMode

Controls how custom-currency cost bases are selected for the subscription.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionChangeCostBasisModeDynamic

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionChangeCostBasisMode("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `BillingSubscriptionChangeCostBasisModeDynamic` | dynamic                                         |
| `BillingSubscriptionChangeCostBasisModePinned`  | pinned                                          |