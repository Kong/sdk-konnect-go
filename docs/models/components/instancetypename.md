# InstanceTypeName

Instance type name to indicate capacity.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.InstanceTypeNameSmall

// Open enum: custom values can be created with a direct type cast
custom := components.InstanceTypeName("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `InstanceTypeNameSmall`  | small                    |
| `InstanceTypeNameMedium` | medium                   |
| `InstanceTypeNameLarge`  | large                    |