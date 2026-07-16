# Behavior

Tax behavior applied to the invoice line item.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BehaviorInclusive

// Open enum: custom values can be created with a direct type cast
custom := components.Behavior("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `BehaviorInclusive` | inclusive           |
| `BehaviorExclusive` | exclusive           |