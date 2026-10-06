# IDTokenParamType

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.IDTokenParamTypeBody

// Open enum: custom values can be created with a direct type cast
custom := components.IDTokenParamType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `IDTokenParamTypeBody`   | body                     |
| `IDTokenParamTypeHeader` | header                   |
| `IDTokenParamTypeQuery`  | query                    |