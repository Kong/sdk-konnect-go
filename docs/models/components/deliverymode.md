# DeliveryMode

Indicates whether the delivery is for a live event or a test.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DeliveryModeLive

// Open enum: custom values can be created with a direct type cast
custom := components.DeliveryMode("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `DeliveryModeLive` | live               |
| `DeliveryModeTest` | test               |