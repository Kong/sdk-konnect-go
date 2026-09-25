# Status

The current validation status of the skill. `pending` means the content has not yet been read and validated.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.StatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.Status("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `StatusPending` | pending         |
| `StatusValid`   | valid           |
| `StatusInvalid` | invalid         |