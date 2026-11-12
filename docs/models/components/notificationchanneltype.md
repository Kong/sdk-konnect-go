# NotificationChannelType

Channel of a user scoped subscription. The http channel is not a user channel: it belongs to org scoped notifications.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.NotificationChannelTypeEmail

// Open enum: custom values can be created with a direct type cast
custom := components.NotificationChannelType("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `NotificationChannelTypeEmail` | EMAIL                          |
| `NotificationChannelTypeInApp` | IN_APP                         |