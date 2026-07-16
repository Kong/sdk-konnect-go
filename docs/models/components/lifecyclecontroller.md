# LifecycleController

Indicates whether the charge lifecycle is controlled by OpenMeter or manually
overridden by the API user.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.LifecycleControllerSystem

// Open enum: custom values can be created with a direct type cast
custom := components.LifecycleController("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `LifecycleControllerSystem` | system                      |
| `LifecycleControllerManual` | manual                      |