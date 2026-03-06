# AwsResourceEndpointConfigState

The current state of the resource config in AWS Resource Endpoint. Possible values:
- `initializing` - The config is in the process of being initialized and is setting up necessary resources.
- `missing` - The config is missing and is no longer accepting new traffic.
- `ready` - The config is fully operational and can route traffic as configured.
- `error` - The config is in an error state, and is not operational.
- `terminating` - The config is in the process of being deleted and is no longer accepting new traffic.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AwsResourceEndpointConfigStateInitializing

// Open enum: custom values can be created with a direct type cast
custom := components.AwsResourceEndpointConfigState("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `AwsResourceEndpointConfigStateInitializing` | initializing                                 |
| `AwsResourceEndpointConfigStateMissing`      | missing                                      |
| `AwsResourceEndpointConfigStateReady`        | ready                                        |
| `AwsResourceEndpointConfigStateError`        | error                                        |
| `AwsResourceEndpointConfigStateTerminating`  | terminating                                  |