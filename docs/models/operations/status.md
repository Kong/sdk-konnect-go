# Status

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/operations"
)

value := operations.StatusAccepted

// Open enum: custom values can be created with a direct type cast
custom := operations.Status("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `StatusAccepted`   | accepted           |
| `StatusInProgress` | in_progress        |
| `StatusCompleted`  | completed          |
| `StatusFailed`     | failed             |