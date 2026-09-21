# Mediation

The mediation type for SASL/PLAIN authentication.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MediationPassthrough

// Open enum: custom values can be created with a direct type cast
custom := components.Mediation("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `MediationPassthrough` | passthrough            |
| `MediationTerminate`   | terminate              |