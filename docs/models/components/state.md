# State

The current state of the managed cache add-on in the data-plane group. Possible values:
- `initializing` - The add-on is in the process of being initialized/updated and is setting up necessary resources for this data-plane group.
- `ready` - The add-on is fully operational for this data-plane group.
- `error` - The add-on is in an error state, and is not operational for this data-plane group.
- `terminating` - The add-on is in the process of being deleted for this data-plane group.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.StateInitializing

// Open enum: custom values can be created with a direct type cast
custom := components.State("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `StateInitializing` | initializing        |
| `StateReady`        | ready               |
| `StateError`        | error               |
| `StateTerminating`  | terminating         |