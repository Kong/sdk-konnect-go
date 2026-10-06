# BillingSubscriptionMigrateResponseStatus

The status of the subscription.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BillingSubscriptionMigrateResponseStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.BillingSubscriptionMigrateResponseStatus("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `BillingSubscriptionMigrateResponseStatusActive`    | active                                              |
| `BillingSubscriptionMigrateResponseStatusInactive`  | inactive                                            |
| `BillingSubscriptionMigrateResponseStatusCanceled`  | canceled                                            |
| `BillingSubscriptionMigrateResponseStatusScheduled` | scheduled                                           |