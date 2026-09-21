# Max

Maximum TLS version to use.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MaxTlSv12

// Open enum: custom values can be created with a direct type cast
custom := components.Max("custom_value")
```


## Values

| Name        | Value       |
| ----------- | ----------- |
| `MaxTlSv12` | TLSv1.2     |
| `MaxTlSv13` | TLSv1.3     |