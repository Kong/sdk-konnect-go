# NotificationStatus

Status of the notification.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.NotificationStatusRead

// Open enum: custom values can be created with a direct type cast
custom := components.NotificationStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `NotificationStatusRead`     | READ                         |
| `NotificationStatusUnread`   | UNREAD                       |
| `NotificationStatusArchived` | ARCHIVED                     |