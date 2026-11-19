# AICostNotificationDeliveryDeliveredStatus

Slack's webhook answered with a `2xx` status, or the email was queued for every recipient at Konnect's email service. A later rejection, suppression, or bounce by the email service doesn't change this status.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostNotificationDeliveryDeliveredStatusDelivered
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `AICostNotificationDeliveryDeliveredStatusDelivered` | delivered                                            |