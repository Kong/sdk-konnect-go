# AIGatewayDataPlaneNodeConfigSyncState

Config sync state.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayDataPlaneNodeConfigSyncStateStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayDataPlaneNodeConfigSyncState("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `AIGatewayDataPlaneNodeConfigSyncStateStateUnspecified` | STATE_UNSPECIFIED                                       |
| `AIGatewayDataPlaneNodeConfigSyncStateStateInSync`      | STATE_IN_SYNC                                           |
| `AIGatewayDataPlaneNodeConfigSyncStateStatePending`     | STATE_PENDING                                           |
| `AIGatewayDataPlaneNodeConfigSyncStateStateResiliency`  | STATE_RESILIENCY                                        |