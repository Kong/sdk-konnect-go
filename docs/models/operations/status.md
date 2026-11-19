# Status

Whether the forecast has predictions. It's `ready` when `predicted` holds them. It's `insufficient_history` when `method` has fewer completed training days than it needs, and `predicted` is then empty.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/operations"
)

value := operations.StatusReady

// Open enum: custom values can be created with a direct type cast
custom := operations.Status("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `StatusReady`               | ready                       |
| `StatusInsufficientHistory` | insufficient_history        |