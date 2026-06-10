# Position

Position and visibility of the payment method reuse agreement.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/metering"
)

value := metering.PositionAuto

// Open enum: custom values can be created with a direct type cast
custom := metering.Position("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `PositionAuto`   | auto             |
| `PositionHidden` | hidden           |