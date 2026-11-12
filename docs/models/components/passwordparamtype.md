# PasswordParamType

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PasswordParamTypeBody

// Open enum: custom values can be created with a direct type cast
custom := components.PasswordParamType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `PasswordParamTypeBody`   | body                      |
| `PasswordParamTypeHeader` | header                    |
| `PasswordParamTypeQuery`  | query                     |