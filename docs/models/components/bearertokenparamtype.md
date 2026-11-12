# BearerTokenParamType

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.BearerTokenParamTypeBody

// Open enum: custom values can be created with a direct type cast
custom := components.BearerTokenParamType("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `BearerTokenParamTypeBody`   | body                         |
| `BearerTokenParamTypeCookie` | cookie                       |
| `BearerTokenParamTypeHeader` | header                       |
| `BearerTokenParamTypeQuery`  | query                        |