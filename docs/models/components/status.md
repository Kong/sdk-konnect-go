# Status

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.StatusSuccess

// Open enum: custom values can be created with a direct type cast
custom := components.Status("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `StatusSuccess` | success         |
| `StatusFailed`  | failed          |