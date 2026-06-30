# BillingSubscriptionStatus

The status of the subscription.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionStatus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `BillingSubscriptionStatusActive`    | active                               |
| `BillingSubscriptionStatusInactive`  | inactive                             |
| `BillingSubscriptionStatusCanceled`  | canceled                             |
| `BillingSubscriptionStatusScheduled` | scheduled                            |