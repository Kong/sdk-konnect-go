# RefreshTokenParamType

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.RefreshTokenParamTypeBody

// Open enum: custom values can be created with a direct type cast
custom := components.RefreshTokenParamType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `RefreshTokenParamTypeBody`   | body                          |
| `RefreshTokenParamTypeHeader` | header                        |
| `RefreshTokenParamTypeQuery`  | query                         |