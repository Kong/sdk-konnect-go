# DeliveryAttemptStatus

Lifecycle state of a delivery attempt.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DeliveryAttemptStatusFailed

// Open enum: custom values can be created with a direct type cast
custom := components.DeliveryAttemptStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `DeliveryAttemptStatusFailed`    | failed                           |
| `DeliveryAttemptStatusSucceeded` | succeeded                        |