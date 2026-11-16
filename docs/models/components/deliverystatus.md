# DeliveryStatus

Delivery status. scheduled means delivery is pending; succeeded means delivery completed successfully; failed means delivery attempts ended without success; expired means the delivery reached its expiry time before it succeeded.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DeliveryStatusScheduled

// Open enum: custom values can be created with a direct type cast
custom := components.DeliveryStatus("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `DeliveryStatusScheduled` | scheduled                 |
| `DeliveryStatusSucceeded` | succeeded                 |
| `DeliveryStatusFailed`    | failed                    |
| `DeliveryStatusExpired`   | expired                   |