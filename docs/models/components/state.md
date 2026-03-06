# State

State of the data-plane group.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.StateCreated

// Open enum: custom values can be created with a direct type cast
custom := components.State("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `StateCreated`      | created             |
| `StateInitializing` | initializing        |
| `StateReady`        | ready               |
| `StateTerminating`  | terminating         |
| `StateTerminated`   | terminated          |