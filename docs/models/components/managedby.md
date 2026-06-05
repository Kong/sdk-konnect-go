# ManagedBy

The charge is managed by the following entity.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ManagedByManual

// Open enum: custom values can be created with a direct type cast
custom := components.ManagedBy("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `ManagedByManual`       | manual                  |
| `ManagedBySystem`       | system                  |
| `ManagedBySubscription` | subscription            |