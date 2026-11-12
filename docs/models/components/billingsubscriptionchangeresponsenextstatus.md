# BillingSubscriptionChangeResponseNextStatus

The status of the subscription.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionChangeResponseNextStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionChangeResponseNextStatus("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `BillingSubscriptionChangeResponseNextStatusActive`    | active                                                 |
| `BillingSubscriptionChangeResponseNextStatusInactive`  | inactive                                               |
| `BillingSubscriptionChangeResponseNextStatusCanceled`  | canceled                                               |
| `BillingSubscriptionChangeResponseNextStatusScheduled` | scheduled                                              |