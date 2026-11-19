# AICostNotificationTestEventDeliveredStatus

Slack's webhook answered the test message with a `2xx` status, or the test email was queued for every recipient at Konnect's email service. A later rejection, suppression, or bounce by the email service doesn't change this status.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostNotificationTestEventDeliveredStatusDelivered
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `AICostNotificationTestEventDeliveredStatusDelivered` | delivered                                             |