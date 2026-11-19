# Posture

How the budget responds when spend reaches its limit.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PostureAdvisory

// Open enum: custom values can be created with a direct type cast
custom := components.Posture("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `PostureAdvisory` | advisory          |
| `PostureEnforce`  | enforce           |