# BillingSubscriptionChangeResponseStatus

The status of the subscription.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionChangeResponseStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionChangeResponseStatus("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `BillingSubscriptionChangeResponseStatusActive`    | active                                             |
| `BillingSubscriptionChangeResponseStatusInactive`  | inactive                                           |
| `BillingSubscriptionChangeResponseStatusCanceled`  | canceled                                           |
| `BillingSubscriptionChangeResponseStatusScheduled` | scheduled                                          |