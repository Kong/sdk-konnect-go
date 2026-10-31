# Shift

The direction to shift surrounding phases to fill the removed phase's span.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ShiftNext

// Open enum: custom values can be created with a direct type cast
custom := components.Shift("custom_value")
```


## Values

| Name        | Value       |
| ----------- | ----------- |
| `ShiftNext` | next        |
| `ShiftPrev` | prev        |