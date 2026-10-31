# BillingSubscriptionMigrateResponseNextStatus

The status of the subscription.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionMigrateResponseNextStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionMigrateResponseNextStatus("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `BillingSubscriptionMigrateResponseNextStatusActive`    | active                                                  |
| `BillingSubscriptionMigrateResponseNextStatusInactive`  | inactive                                                |
| `BillingSubscriptionMigrateResponseNextStatusCanceled`  | canceled                                                |
| `BillingSubscriptionMigrateResponseNextStatusScheduled` | scheduled                                               |