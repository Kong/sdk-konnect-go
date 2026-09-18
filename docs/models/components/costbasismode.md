# CostBasisMode

Controls how custom-currency cost bases are selected for the subscription.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CostBasisModeDynamic

// Open enum: custom values can be created with a direct type cast
custom := components.CostBasisMode("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `CostBasisModeDynamic` | dynamic                |
| `CostBasisModePinned`  | pinned                 |