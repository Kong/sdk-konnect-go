# Capabilities

**`skills` requires a minimum runtime version of `2.2`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CapabilitiesBatches

// Open enum: custom values can be created with a direct type cast
custom := components.Capabilities("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `CapabilitiesBatches` | batches               |
| `CapabilitiesFiles`   | files                 |
| `CapabilitiesSkills`  | skills                |