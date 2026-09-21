# SubscriptionStatus

The status of the subscription.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.SubscriptionStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.SubscriptionStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `SubscriptionStatusActive`    | active                        |
| `SubscriptionStatusInactive`  | inactive                      |
| `SubscriptionStatusCanceled`  | canceled                      |
| `SubscriptionStatusScheduled` | scheduled                     |