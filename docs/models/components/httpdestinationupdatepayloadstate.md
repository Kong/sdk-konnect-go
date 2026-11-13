# HTTPDestinationUpdatePayloadState

Set to enabled to allow notifications to be sent, or disabled to pause them. Setting a destination that was automatically disabled back to enabled restarts delivery attempts.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.HTTPDestinationUpdatePayloadStateEnabled

// Open enum: custom values can be created with a direct type cast
custom := components.HTTPDestinationUpdatePayloadState("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `HTTPDestinationUpdatePayloadStateEnabled`  | enabled                                     |
| `HTTPDestinationUpdatePayloadStateDisabled` | disabled                                    |