# DestinationState

Destination status. pending_verification means the destination is waiting for verification; enabled means it can receive notifications; disabled means it is paused; disabled_system means it was disabled automatically.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DestinationStatePendingVerification

// Open enum: custom values can be created with a direct type cast
custom := components.DestinationState("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `DestinationStatePendingVerification` | pending_verification                  |
| `DestinationStateEnabled`             | enabled                               |
| `DestinationStateDisabled`            | disabled                              |
| `DestinationStateDisabledSystem`      | disabled_system                       |