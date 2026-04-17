# CustomDomainState

State of the custom domain.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CustomDomainStateCreated

// Open enum: custom values can be created with a direct type cast
custom := components.CustomDomainState("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `CustomDomainStateCreated`      | created                         |
| `CustomDomainStateInitializing` | initializing                    |
| `CustomDomainStateReady`        | ready                           |
| `CustomDomainStateTerminating`  | terminating                     |
| `CustomDomainStateTerminated`   | terminated                      |
| `CustomDomainStateError`        | error                           |