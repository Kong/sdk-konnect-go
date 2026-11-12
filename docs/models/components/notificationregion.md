# NotificationRegion

Region associated to a notification.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.NotificationRegionUs

// Open enum: custom values can be created with a direct type cast
custom := components.NotificationRegion("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `NotificationRegionUs`       | US                           |
| `NotificationRegionEu`       | EU                           |
| `NotificationRegionAu`       | AU                           |
| `NotificationRegionMe`       | ME                           |
| `NotificationRegionIn`       | IN                           |
| `NotificationRegionWildcard` | *                            |