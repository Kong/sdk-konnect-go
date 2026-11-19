# AICostNotificationDeliveryStatus

The progress of one message to one channel:

- `pending`: the message is queued, being sent, or waiting for a retry.
- `delivered`: Slack's webhook answered with a `2xx` status, or the email was queued for every recipient at Konnect's email service. A later rejection, suppression, or bounce by the email service doesn't change this status.
- `failed`: the message couldn't be delivered, and `error` says why.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostNotificationDeliveryStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.AICostNotificationDeliveryStatus("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `AICostNotificationDeliveryStatusPending`   | pending                                     |
| `AICostNotificationDeliveryStatusDelivered` | delivered                                   |
| `AICostNotificationDeliveryStatusFailed`    | failed                                      |