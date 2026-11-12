# ClientCredentialsParamType

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ClientCredentialsParamTypeBody

// Open enum: custom values can be created with a direct type cast
custom := components.ClientCredentialsParamType("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `ClientCredentialsParamTypeBody`   | body                               |
| `ClientCredentialsParamTypeHeader` | header                             |
| `ClientCredentialsParamTypeQuery`  | query                              |