# CPUProfileMode

The mode of cpu profile

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CPUProfileModeTime

// Open enum: custom values can be created with a direct type cast
custom := components.CPUProfileMode("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `CPUProfileModeTime`        | time                        |
| `CPUProfileModeInstruction` | instruction                 |