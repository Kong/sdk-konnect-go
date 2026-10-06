# BillingSubscriptionSubscriptionStatus

The status of the subscription.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionSubscriptionStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionSubscriptionStatus("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `BillingSubscriptionSubscriptionStatusActive`    | active                                           |
| `BillingSubscriptionSubscriptionStatusInactive`  | inactive                                         |
| `BillingSubscriptionSubscriptionStatusCanceled`  | canceled                                         |
| `BillingSubscriptionSubscriptionStatusScheduled` | scheduled                                        |